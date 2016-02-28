---
layout: post
title: Create Random Safaricom Kenya Phone Numbers using Faker regexify
author: Fred Muya
---

[Faker](https://github.com/fzaninotto/Faker){:target="_blank"} is an amazing PHP library you can use to generate fake data, e.g names, addresses, etc.

One of it's amazing features is the ability to create random data based on a regular expression. Using the `regexify` function, you can create a random value that satisfies a given regular expression.

A sample use case is generating test phone numbers based on a phone number regular expression. In my case, I have some test cases that require me to generate a Safaricom Kenya mobile number.

The regex I'm using for the phone number is: `/(\+?254|0){1}[7]{1}([0-2]{1}[0-9]{1}|[9]{1}[0-2]{1})[0-9]{6}/`. That matches the following:

* Phone numbers starting with '0' (national format), or '254' (international format)
* Followed by '7[0-2][0-9]' or '7[9][0-2]', e.g. 701, 711, 712, 790, 791, 792, etc
* Followed by 6 digits in the range of 0-9

NB: For phone number validation purposes, the regex could be made more forgiving, e.g. by allowing a space between the '254' and the rest of the number, etc

Here's an example of how I'm using it:

{% highlight bash %}

# using Laravel's artisan tinker for command line here, the code can be used directly in your PHP code
>>> use Faker\Factory as Faker
=> false
>>> $faker = Faker::create()
>>> $regex = '/(\+?254|0){1}[7]{1}([0-2]{1}[0-9]{1}|[9]{1}[0-2]{1})[0-9]{6}/'
=> "/(\+?254|0){1}[7]{1}([0-2]{1}[0-9]{1}|[9]{1}[0-2]{1})[0-9]{6}/"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "254791170875"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "254792300310"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "0700490004"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "+254791100914"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "+254705774077"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "0720958391"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "+254719424565"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "+254705699957"
>>> $samplePhoneNumber = $faker->regexify($regex)
=> "254726015158"

{% endhighlight %}

In this case, a random valid valid Safaricom Kenya phone number is generated every time you use the function. Feel free to use it in your code!

I've created a gist with the pure PHP code [here](https://gist.github.com/muya/d44dd3d379bcdef175cf){:target="_blank"}

PS: I have a list of regex you can use to validate Kenyan phone numbers [here](https://gist.github.com/muya/ce5a18a3f119cc4ac286){:target="_blank"}, feel free to use that too. Additions/modifications are welcome ;-)
