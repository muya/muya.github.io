---
layout: post
title: Register Custom Doctrine Column Type Mappings for Laravel Backpack CRUD
author: Fred Muya
---

I've been using [Laravel Backpack CRUD](https://laravel-backpack.readme.io/docs) to build the backend for a personal application, and I must say - great tool! Generation of the CRUD part of the application is only an artisan command away! Reminds me of the Yii Auto generator.

However, I run into an issue with one of my database columns that has type `JSON`. After running the `backpack:crud` command, the UI for it wasn't loading. Instead the error below was showing:

{% highlight text %}
2017-01-02 09:28:12 [ERROR] (local): Doctrine\DBAL\DBALException: Unknown database type json requested, Doctrine\DBAL\Platforms\MySQL57Platform may not support it. in /path/to/laravel/app/vendor/doctrine/dbal/lib/Doctrine/DBAL/Platforms/AbstractPlatform.php:423
$
{% endhighlight %}

Upon doing some research (read extensive Googling/Github Issue checking), it turns out it's a DBAL-related issue, affecting specific column types in MySQL (including `enum` & `json`-like columns).

For more details on why `enum`, in particular, has some issues: [http://docs.doctrine-project.org/projects/doctrine-orm/en/latest/cookbook/mysql-enums.html](http://docs.doctrine-project.org/projects/doctrine-orm/en/latest/cookbook/mysql-enums.html)

I came across a suggestion for the fix in the Laravel Backpack CRUD here: [Register Custom Doctrine Column Type Mappings](https://github.com/Laravel-Backpack/CRUD/issues/269) - and I've detailed how I implemented the fix in my application.

My fix involves:

- Implementing a modified version of the CrudController, which;
- Implements a modified CrudPanel class that includes addition of the custom column types mentioned in the fix

First, create a directory called `CRUD` under the app directory of your Laravel application (you can call it anything you please, be sure to update the code below to match your naming), then create a new PHP file there which will have the modified CrudPanel class:
{% highlight bash %}

$ cd /path/to/laravel/app
$ mkdir app/CRUD
$ touch app/CRUD/CustomCrudPanel.php

# your path should look something like:
$ tree app/
app/
├── Console
│   ├── Commands
│   │   └── Inspire.php
│   └── Kernel.php
├── CRUD
│   └── CustomCrudPanel.php
├── Events
│   └── Event.php
├── Exceptions
│   └── Handler.php
├── Http
│   ├── Controllers
...
{% endhighlight %}

Next, create a custom CrudController file which will have our extended & modified CrudController class. I put mine in `app/Http/Controllers/Admin/` - where the other Backpack CRUD controllers are.
{% highlight bash %}

$ cd /path/to/laravel/app
$ touch app/Http/Controllers/Admin/CustomCrudController.php

{% endhighlight %}

Put the code below in `CustomCrudPanel.php`:
{% highlight php %}
<?php

// update this to match your app's namespace
namespace MyApp\CRUD;

use Backpack\CRUD\CrudPanel;
use Backpack\CRUD\PanelTraits\AutoSet;

class CustomCrudPanel extends CrudPanel
{
    use AutoSet {
        // we're overriding the original method from the AutoSet trait
        setFromDb as parentSetFromDb;
    }

    public function setFromDb()
    {
        // register custom column types
        $this->addCustomDoctrineColumnTypes();

        // call the parent method so that all attributes are initialized properly
        $this->parentSetFromDb();
    }

    // this is the fix suggested from the Github Issue (https://goo.gl/wN7dEd) Thank you @zschuessler
    public function addCustomDoctrineColumnTypes()
    {
        $dbPlatform = \Schema::getConnection()->getDoctrineSchemaManager()->getDatabasePlatform();
        $dbPlatform->registerDoctrineTypeMapping('enum', 'string');
        $dbPlatform->registerDoctrineTypeMapping('json', 'json_array');
    }
}

{% endhighlight %}

Put the code below in `CustomCrudController.php`:
{% highlight php %}
<?php

// update this to match your app's namespace
namespace MyApp\Http\Controllers\Admin;

use Backpack\CRUD\app\Http\Controllers\CrudController;
use MyApp\CRUD\CustomCrudPanel as CrudPanel;

class CustomCrudController extends CrudController
{

    public function __construct()
    {
        if (!$this->crud) {
            $this->crud = new CrudPanel();

            // call the setup function inside this closure to also have the request there
            // this way, developers can use things stored in session (auth variables, etc)
            $this->middleware(function ($request, $next) {
                $this->request = $request;
                $this->crud->request = $request;
                $this->setup();

                return $next($request);
            });
        }
    }
}
{% endhighlight %}

Finally, in your auto-generated `CrudController` class for the model with an `enum` or `json` column, update it to extend the CustomCrudController class instead.

E.g. if the `CrudController` class with the `json` column is called `PlaylistCrudController`, it should look something like:
{% highlight php %}
<?php

// update this to match your app's namespace
namespace MyApp\Http\Controllers\Admin;

// use the CustomCrudController class instead of Backpack\CRUD\app\Http\Controllers\CrudController
use MyApp\Http\Controllers\Admin\CustomCrudController as CrudController;
use MyApp\Http\Requests\PlaylistRequest as StoreRequest;
use MyApp\Http\Requests\PlaylistRequest as UpdateRequest;

class PlaylistCrudController extends CrudController
{

    public function setUp()
    {
        $this->crud->setModel('MyApp\Models\Playlist');
        $this->crud->setRoute('admin/playlist');
        $this->crud->setEntityNameStrings('playlist', 'playlist');

        $this->crud->setFromDb();
    }

    public function store(StoreRequest $request)
    {
        // your additional operations before save here
        $redirect_location = parent::storeCrud();

        return $redirect_location;
    }

    public function update(UpdateRequest $request)
    {
        // your additional operations before save here
        $redirect_location = parent::updateCrud();

        return $redirect_location;
    }
}
{% endhighlight %}

If all goes well, you should be able to view the CRUD for our troublesome model.

Credits for the suggested fix go to [@zschuessler](https://github.com/zschuessler) from [this Github Issue](https://github.com/Laravel-Backpack/CRUD/issues/269) on the Laravel Backpack CRUD.

I also ended up learning a lot about PHP Traits that I didn't know before.

Thank you Laravel Backpack CRUD project for such an amazing tool!

If you have any comments/suggestions/improvements, feel free to reach out in the comments.

Happy Coding! (And all the best this year?!)
