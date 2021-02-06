---
layout: post
title: Using Jinja Template Filters in Ansible
date: '2021-02-06'
author: Fred Muya
excerpt: How to pass additional arguments to Jinja Template filters in Ansible
tags:
- jinja
- ansible
- jinja template
- trim
- strip characters
- int jinja filter
- jinja filter
---

[Jinja](https://jinja.palletsprojects.com/en/2.11.x/) is a templating language that is built for Python.

Ansible has support for using Jinja Templating in its playbooks ([Learn More](https://docs.ansible.com/ansible/latest/user_guide/playbooks_templating.html)).

_**Note**: All the code snippets used in this post are [available on GitHub](https://github.com/muya/muya.github.io/blob/master/snippets/2021-02-06/ansible-playbook-demo-jinja-filters-in-ansible.yml)_


While working with some of the filters that have an option to pass in additional arguments, it wasn't very clear to me how to pass in the arguments within an Ansible playbook.

For example, I wanted to use the `trim` filter to remove specific characters from a string I was working with.

{% highlight yaml %}

# Original string: *.helloworld.com
# Required string: helloworld.com


{% endhighlight %}

By default, the trim filter strips leading and trailing whitespace from the string provided.

{% highlight yaml %}

# word_with_whitespace = "    helloworld   "
- name: "Show default functionality of trim filter"
ansible.builtin.debug:
  msg: "Before: [{% raw %}{{ word_with_whitespace }}{% endraw %}] | After: [{% raw %}{{ word_with_whitespace | trim }}{% endraw %}]"

{% endhighlight %}

{% highlight plaintext %}

TASK [Show default functionality of trim filter]
ok: [localhost] => {
    "changed": false,
    "msg": "Before: [    helloworld   ] | After: [helloworld]"
}

{% endhighlight %}


Now, let's say you want to trim the `"*."` characters from our example above.

According to the [Jinja `trim` documentation](https://jinja.palletsprojects.com/en/2.11.x/templates/#trim), the signature for the method is as follows:

{% highlight plaintext %}

trim(value, chars=None)

{% endhighlight %}

###### But what does this translate to when using it in Ansible?

It turns out that when working with the Jinja filters, the value to which you're applying the filter is automatically passed on to the filter as the first argument.

Therefore, the following snippet in Ansible:

{% highlight plaintext %}


{% raw %}{{ word_with_whitespace | trim }}{% endraw %}


{% endhighlight %}

gets translated to this (in Python):

{% highlight python %}


trim(word_with_whitespace)


{% endhighlight %}


So if we want to trim `"*."` from our string, we need to ensure that the Python code that gets executed is:

{% highlight python %}


trim(word_with_whitespace, "*.")


{% endhighlight %}


In Ansible, this becomes:

{% highlight plaintext %}


{% raw %}{{ word_with_whitespace | trim("*.") }}{% endraw %}


{% endhighlight %}


So the snippet to trim this from our string in a playbook becomes:

{% highlight yaml %}

# word_with_other_chars = "*.helloworld.com"
- name: Show trim filter with additional arguments
ansible.builtin.debug:
  msg: "Before: [{% raw %}{{ word_with_other_chars }}{% endraw %}] | After: [{% raw %}{{ word_with_other_chars | trim('*.') }}{% endraw %}]"


{% endhighlight %}

which yields:

{% highlight plaintext %}

TASK [Show trim filter with additional arguments]
ok: [localhost] => {
    "changed": false,
    "msg": "Before: [*.helloworld.com] | After: [helloworld.com]"
}

{% endhighlight %}


The same principle applies to other Jinja filters; for example, the [`int()` filter](https://jinja.palletsprojects.com/en/2.11.x/templates/#int), which converts a given value to an integer.

###### int Jinja Filter Method Signature

{% highlight plaintext %}

trim(value, chars=None)

{% endhighlight %}


###### Ansible Task Snippets

{% highlight yaml %}

# str_value_to_convert_to_int: "5278"
# hex_value_to_convert_to_int: "0x149E"
- name: Show trim filter with additional arguments
  - name: Show int filter without arguments
    ansible.builtin.debug:
      msg: "Original Value: [{% raw %}{{ str_value_to_convert_to_int }}{% endraw %}] | Converted to int: [{% raw %}{{ str_value_to_convert_to_int | int }}{% endraw %}]"

  - name: Show int filter with arguments (convert string from hexadecimal (base 16) to decimal (base 10))
    ansible.builtin.debug:
      msg: "Original Value: [{% raw %}{{ hex_value_to_convert_to_int }}{% endraw %}] | Converted to int: [{% raw %}{{ hex_value_to_convert_to_int | int(0, 16) }}{% endraw %}]"

{% endhighlight %}

###### Ansible Output

{% highlight plaintext %}

TASK [Show int filter without arguments]
ok: [localhost] => {
    "changed": false,
    "msg": "Original Value: [5278] | Converted to int: [5278]"
}

TASK [Show int filter with arguments (convert string from hexadecimal (base 16) to decimal (base 10))]
ok: [localhost] => {
    "changed": false,
    "msg": "Original Value: [0x149E] | Converted to int: [5278]"
}

{% endhighlight %}


In this blog post, we demonstrated how to use Jinja filters that have multiple arguments, within Ansible.

Happy coding, and stay safe!
