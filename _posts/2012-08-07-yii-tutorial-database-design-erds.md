---
layout: post
title: 'Yii Tutorial: Database Design & ERDs'
date: '2012-08-07T07:26:00.001+03:00'
author: Fred Muya
tags:
- Yii Framework
modified_time: '2012-08-10T08:34:57.384+03:00'
thumbnail: http://4.bp.blogspot.com/-Bb9iZpDp6UY/UCCX2gZxuhI/AAAAAAAAAQo/y87ckhIEk8A/s72-c/StudentPortal+-+ERD.png
blogger_id: tag:blogger.com,1999:blog-3630098189866392828.post-2482110224443790786
blogger_orig_url: http://encore254.blogspot.com/2012/08/yii-tutorial-database-design-erds.html
---

T
This is the second part of the Yii tutorial. It covers the database design as well as the ERD for the project we are going to embark on.

The diagram below shows a basic ERD of the system. Only the major columns in the tables are shown in the diagram. There are other columns that are common to
all the tables:

- status(tinyint[3]) - describes the record’s entity state
- dateCreated(datetime) - describes the date the record was inserted
- insertedBy(int[11]) - the userID of the user who inserted the record
- dateModified(datetime) - the date the record was last modified/updated
- updatedBy(int[11]) - the userID of the user who last updated it

![Student Portal Database ERD]({{ site.url }}/images/2012-08-07/StudentPortal-ERD.png)

More information about the table fields will be given as we use them.

You can download the [SQL file](https://github.com/muya/student-portal/raw/master/db/studentPortal_28052013_1322.sql) needed to create the database.
