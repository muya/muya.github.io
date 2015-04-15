---
layout: post
title: Implementing Timestamped Models in PeeWee
author: Fred Muya
---

`date_created` and `date_updated` (or variations of the same) are important fields for any table. `date_created` to show when a record was created, `date_updated` to show when it was last updated. They are invaluable, especially for audit logs.

I've been using [PeeWee](https://github.com/coleifer/peewee "PeeWee Github page"), a lightweight Python ORM, for a small project at work. Since I'm using SQLite3 for the project, I decided to use the ORM to ease the database access work. This post will show how I implemented a Timestamped Model.

The rules for my timestamped model are:

- Set `date_created` field to current timestamp any time a record is **created**
- Set `date_updated` field to current timestamp any time a record is **updated**

I have a `BaseModel` from which all my PeeWee models inherit

{% highlight python %}
database = SqliteDatabase(Config.get("DATABASE"), **{})

class BaseModel(Model):
    class Meta:
        database = database
{% endhighlight %}

Then my `TimestampedModel` inherits from the BaseModel:

{% highlight python %}
class TimestampedModel(BaseModel):
    def save(self, *args, **kwargs):
        if self._get_pk_value() is None:
            # this is a create operation, set the date_created field
            self.date_created = datetime.datetime.now().strftime(
                "%Y-%m-%d %H:%M:%S")
        self.date_updated = datetime.datetime.now().strftime(
            "%Y-%m-%d %H:%M:%S")
        return super(TimestampedModel, self).save(*args, **kwargs)
{% endhighlight %}

In the snippet above, I've overridden the `save` method to do my "timestamping". I found the `_get_pk_value()` in the source code (go open source!).

Then in my model, I just inherit from the `TimestampedModel`:

{% highlight python %}
class Person(TimestampedModel):
    date_created = DateTimeField()
    date_updated = DateTimeField()

    # other cool fields
{% endhighlight %}

An alternative, especially for the `date_created` field, would have been to have a default value in the model, i.e. something like:

{% highlight python %}
class Person(BaseModel):
    date_created = DateTimeField(default=datetime.datetime.now().strftime(
        "%Y-%m-%d %H:%M:%S"))
{% endhighlight %}

However, I noticed that for create operations that were close together, e.g. a few seconds apart, they were being given the same `date_created` value. I'm not sure if I was doing something wrong, but the solution above resolved this issue for me.

Of course, using the SQL triggers (e.g. `ON CREATE CURRENT_TIMESTAMP`) would also work (especially for MySQL). For SQLite however, you run into issues when writing tests and you want to use the memory for temporary storage. You might get value missing errors for columns that are `NOT NULL`, since the SQL schema isn't used when accessing the database from within the tests.

I hope this helps someone out there! All the best!

-- Back to code!
