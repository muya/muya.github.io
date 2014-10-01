---
layout: post
title: Using Gii (Part 2)
date: '2012-08-23T08:00:00.001+03:00'
author: Fred Muya
excerpt: Part 2 of the Using Gii Tutorial
tags:
- Yii Framework
modified_time: '2012-10-23T08:20:30.676+03:00'
---

**NB:** Please note that if you are unable to access Gii with the url given in the previous post, you may have to change some configuration in `main.php`. Uncomment the part that describes the `urlManager` property. This allows us to use more 'user-friendly' URLs such as `http://{YOUR-SERVER}/StudentPortal/index.php/gii/model` rather than `http://{YOUR-SERVER}/StudentPortal/index.php?r=gii/model`. The formatting of the URL will be explained later in this post.

In this section, we are going to create our first model, controller, and its views. Since StudentPortal uses Role-Based Access Control, it makes sense to start with 'Users', as they are core to the system. Having successfully accessed the Gii user interface in [Part 1](http://encore254.blogspot.com/2012/08/using-gii-part-1.html), we are going to create the MVCs.

First step, go to [http://{YOUR-SERVER}/StudentPortal/index.php/gii](http://{YOUR-SERVER}/StudentPortal/index.php/gii). On the side-menu, there is a list of options that allow you to generate:

- Controller
- CRUD
- Form
- Model
- Module

The controller generator allows you to generate controller classes, which extend Yii’s Controller class. The Crud Generator allows you to generate the views and controllers for a certain model. The Form Generator generates a view with a form to input data, for a specific model. The Model Generator generates models for a specific table, which extend Yii’s CActiveRecord Class. Finally, the Module Controller generates modules for your system, should you need to have separate modules (e.g. UI, Apps, APIs, etc).

For now, we’ll be using the Model and Crud Generators.

#### Model Generator
Click on Model Generator. The form displayed has several fields:

- Table prefix - if the names of the tables in your database start with a certain prefix, such as 'tbl_', Yii will automatically append this prefix when you are referencing the tables in your application, and also Gii will remove it when making the models. Our tables do not use this convention, so we will ignore this field.
- Table name - the name of the table whose model you are generating. This is case-sensitive, so you have to specify the table name as used in the database. If you had provided a prefix, you don’t have to give the full name of the table in this field. For example, if 'tbl_' was the prefix, so a table name would be something like 'tbl_users', you only have to type 'users' in this field. Also note, that as you start typing, a list of matching table names may show up in a list for you to select from, these table names are loaded from the database specified in your config file.
- Base Class - this is the class from which the generated model class will extend. Default is Yii’s CActiveRecord class.
- Model Class - The name of the model class you want to generate. Generally, by OOP standards, class names start with an upper-case letter. Therefore Gii will suggest a name for your model class, usually the table name with an upper-case first letter. For example, if the table name is 'users', the model class name 'Users' will be suggested. It is advisable to stick to this convention.
- Model Path - this describes the location where the generated model file will be generated. By default, the path is `application.models`. This means that the file will be put under `protected/models/`. We’ll leave this as it is.
- Build Relations (check box) - this, if checked creates relations, if they are defined in the specified table. Relations refer to the relationships defined between the tables. If foreign key constraints have been defined in the database, these relationships will also be defined in the model. _**Note**: these relations can also be defined manually_.
- Code Template - this is the path to where the template used to generate the model is. This is important especially in a scenario where you need to modify all your models to have a different format from what Gii recommends. This will come in handy later in the development of our application.

For the StudentPortal application:

1.  Leave the Table Prefix field blank
2.  Enter 'users' (without the quotes) as the table name
3.  Leave Base Class and Model Path as they are
4.  Ensure Build Relations is checked
5.  Leave the code template as it is

On clicking 'Preview', a small table appears at the bottom, one column with the code file path, the other with a check box. Items in the first column describe the files that will be generated, the check box on the right column specifies if the corresponding file will be generated. Note that at this point, the files haven’t been created yet. You can preview contents of a specific file by clicking on the file name.

Since we are (hopefully) satisfied with what we have, we can go ahead and generate the file, by clicking 'Generate'. If you now look in `protected/models/`, you will find a file named `Users.php`. This is the model for the users table.
**NB:** If you get an error saying that the code could not be generated, ensure that your web server has full write permissions to the application folder. In a development environment (Unix-based systems), you can do:
{% highlight bash %}
$ sudo chmod o+rw -R *
{% endhighlight %}
in the application folder. For windows, check out how to do that [here](http://answers.microsoft.com/en-us/windows/forum/windows_vista-files/how-do-i-change-folder-and-file-permissions/465f2b42-63dd-4486-8dd1-c870290efeed){:target="_blank"}.

#### Crud Generator
Click on Crud Generator. The form displayed has several fields:

- Model Class - the name of the Model class whose views and controller will be generated
- Controller ID - the ID of the controller that will be generated for that model. Note that the Crud Generator generates both a controller and the views.
- Base Controller Class - the Class from which the Controller will be extended. The default is Yii’s Controller class
- Code Template - this is the path to where the templates used to generate the views and controller are

For the StudentPortal application:

1. Enter the model class name - 'Users'
2. Leave the other fields as they are (The controller ID field will be automatically populated)
3. Click Preview.
4. Click Generate.

The controller and views will be under `protected/controllers` and `protected/views/users` respectively.

We have now created the Model, Controller and Views for users. You can check out what you’ve created so far by going to [http://{YOUR-SERVER}/StudentPortal/index.php/users](http://localhost/StudentPortal/index.php/users){:target="_blank"}
If you are redirected to a log in page, log in using `admin/admin` as the username/password.

_**NB:** About the URL manager... this allows us to have more readable URLs. Basically, the first part after index.php, refers to the controller, the second part to the controller action. If there are other parts to the URL, they refer to parameters that are passed to the controller action. We’ll learn more about what controller actions are in a later post._

In the next post, we'll be dissecting our magically created code
