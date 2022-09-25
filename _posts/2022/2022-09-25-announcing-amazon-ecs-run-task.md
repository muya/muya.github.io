---
layout: post
title: 'Announcing AWS ECS Run Task GitHub Action'
date: '2022-09-25'
author: Fred Muya
excerpt: 'A GitHub Action to trigger your ECS Task: <a href="https://github.com/muya/amazon-ecs-run-task">muya/amazon-ecs-run-task</a>'
tags:
- aws
- ecs
- aws-ecs
- fargate
- ecs-fargate
- fargate-spot
- ecs-fargate-spot
- github-actions
- ecs-task
---

[¡Sawubona!](https://translate.google.com/?sl=zu&tl=en&text=sawubona&op=translate) Welcome to another blog post!

_Side Note: [Zulu is finally available on Duolingo](https://blog.duolingo.com/welcome-zulu-to-the-language-family/), and I've been learning a little. It's great to see more African languages being added and made available to learn._

Today, we're announcing the release of [muya/amazon-ecs-run-task](https://github.com/muya/amazon-ecs-run-task), a GitHub Action that enables you to trigger your AWS Elastic Container Service (ECS) tasks.

## What Are ECS Tasks?

ECS allows one to deploy a container using several options:

- [Service](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_services.html) - with this option, an instance gets spun up, and runs continuously until explicitly stopped. It's ideal for applications that need to be running all the time, like a web application, API, background task, etc. With this option, ECS will always ensure that the desired number of services are running at any given time, with the items getting replaced in case they exit.
- [Task](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs_run_task.html) - with this option, the instance will run and then exit once it's done with its work. This is ideal for one-off tasks such as running database migrations (which was my use-case), batch jobs that run periodically, etc.

## Why muya/amazon-ecs-run-task GitHub Action?

AWS provides an official GitHub Action: [aws-actions/amazon-ecs-deploy-task-definition](https://github.com/aws-actions/amazon-ecs-deploy-task-definition), but its configuration options only allow one to deploy a "service".

This means that any services deployed using this action will result in a continuously running deployment.


> *Side Note:* I learnt this the hard way 😅. I thought I could work around this by having my application start up, execute a task, and then gracefully exit.
> While the application did this well, ECS would keep spinning it up again once it exited, since Services need to maintain a certain number of instances at any given time.

This brings us to other GitHub Actions which enable deployment of tasks.

They do this by:
- Branching off the original [aws-actions/amazon-ecs-deploy-task-definition](https://github.com/aws-actions/amazon-ecs-deploy-task-definition) project
- Updating it to call the `ecs.runTask` method that's available in the [AWS ECS SDK](https://docs.aws.amazon.com/AWSJavaScriptSDK/latest/AWS/ECS.html#runTask-property)

The one that's seems most popular is [@smitp's amazon-ecs-run-task](https://github.com/smitp/amazon-ecs-run-task).

In addition to enabling the `runTask` functionality, it provides options to:

- Set a desired count for instances that should be running (using `"count"` option)
- Set the user who started the job (using `"started-by"` option)
- Set a `"wait-for-finish"` option, which will cause GitHub Actions to wait for the task to be stopped before proceeding to the next job (similar to `"wait-for-service-stability"` option in the original)

However, it lacked options for specifying the [Launch Type](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html), and / or a [capacity provider strategy](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html).

Additionally, the repository seems not to have been updated in a while.

Thankfully, this PR [#76 Add support for capacity provider strategy attribute](https://github.com/smitp/amazon-ecs-run-task/pull/76) by [evgeniy-b](https://github.com/evgeniy-b) proposed a way to set the capacity provider strategy.

The approach taken allows setting of a JSON representation of a capacity provider strategy, such as:

{% highlight json %}

[{"capacityProvider": "FARGATE_SPOT", "base": 0, "weight": 1}]

{% endhighlight %}

See the [AWS Capacity Provider Concepts Documentation](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cluster-capacity-providers.html) for more details about what options that can be set.

With this, we can configure usage of `FARGATE`, `FARGATE_SPOT` or whatever you have for your setup. This is useful, since using `FARGATE_SPOT` may end up being [up to 70% cheaper than `FARGATE`](https://aws.amazon.com/fargate/pricing/) if the use case fits your needs.

Then finally, we have the version by [@YashdalfTheGray](https://github.com/YashdalfTheGray/amazon-ecs-run-task/) (forked off smitp/amazon-ecs-run-task), which added functionality to set:
- `launch-type`
- options for the network configuration required by `runTask` (`subnets`, `security-groups`, `assign-public-ip`).

The network configuration options are particularly useful if you're setting these up ECS within a private VPC.


## Putting it All Together

[muya/amazon-ecs-run-task](https://github.com/muya/amazon-ecs-run-task) brings together all of the above mentioned functionality into one action!

This means that we can set:

- `launch-type` OR `capacity-provider-strategy` - to specify how the deployment should happen. (Note: one can set either the launch type, or the capacity strategy, but not both. This action will use `capacity-provider-strategy` if both are set).
- Network Configuration for the `awsvpcConfiguration` option:
    - `subnets` - a comma separated list of subnets to set.
    - `security-groups` - a comma separated list of security groups to set.
    - `assign-public-ip` - either `ENABLED` or `DISABLED`. Defaults to `DISABLED`.

These options allow fine-tuning of the deployment to suit your needs!

### Examples

The examples below assume that the task definitions have been set up accordingly. The original [aws-actions/amazon-ecs-deploy-task-definition](https://github.com/aws-actions/amazon-ecs-deploy-task-definition) action has this documented.

#### Deploy ECS Task Using FARGATE

The example below demonstrates a job that will deploy using `FARGATE`. This will result in the default capacity provider strategy being used.

{% highlight yaml %}

      - name: Run as Amazon ECS Task
        uses: muya/amazon-ecs-run-task@v1
        with:
          # This example assumes task definition got set by previous job. Adjust as needed
          task-definition: {% raw %}${{ steps.migrations-task-def.outputs.task-definition }}{% endraw %}
          count: 1
          started-by: {% raw %}github-actions-${{ github.actor }}{% endraw %}
          cluster: 'your-sparkling-cluster'
          launch-type: FARGATE
          subnets: subnet-09e2280498bf, subnet-069c3dbe9116
          security-groups: sg-a230bf8e0fe5, sg-fe89565d2b14, sg-d874051fc1cc
          assign-public-ip: ENABLED
          # Wait for the task to exit on ECS before progressing
          wait-for-finish: true
          # How long we should allow the job to run on GitHub Actions
          wait-for-minutes: 10

{% endhighlight %}



#### Deploy ECS Task Using FARGATE_SPOT

The example below demonstrates a job that will deploy using `FARGATE_SPOT`.

{% highlight yaml %}

      - name: Run as Amazon ECS Task
        uses: muya/amazon-ecs-run-task@v1
        with:
          # This example assumes task definition got set by previous job. Adjust as needed
          task-definition: {% raw %}${{ steps.migrations-task-def.outputs.task-definition }}{% endraw %}
          count: 1
          started-by: {% raw %}github-actions-${{ github.actor }}{% endraw %}
          cluster: 'a-cruising-cluster'
          subnets: subnet-09e2280498bf, subnet-069c3dbe9116
          security-groups: sg-a230bf8e0fe5, sg-fe89565d2b14, sg-d874051fc1cc
          assign-public-ip: ENABLED
          # Wait for the task to exit on ECS before progressing
          wait-for-finish: true
          # How long we should allow the job to run on GitHub Actions
          wait-for-minutes: 10
          capacity-provider-strategy: '[{"capacityProvider": "FARGATE_SPOT", "base": 0, "weight": 1}]'

{% endhighlight %}


## Conclusion

And that's it!

[muya/amazon-ecs-run-task](https://github.com/muya/amazon-ecs-run-task) brings together amazing additions from different contributors to make deploying your ECS tasks from GitHub Actions much easier!

I hope you find it useful!

Until next time, happy coding!
