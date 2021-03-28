---
layout: post
title: Parsing JSON Data Objects having Number-like Names in SnowFlake
date: '2021-01-02'
author: Fred Muya
excerpt: Parse number-like JSON names in SnowFlake
tags:
- snowflake
- parse
- json
- parse-json
- bracket-notation
---

[SnowFlake](https://www.snowflake.com/cloud-data-platform/) is a cloud data solution that provides several tools to allow easier access to your data. According to their website, it's their solution for data warehousing, data lakes, data engineering, data science, data application development, and for securely sharing and consuming shared data.

Sometimes, due to factors outside of our control, it may be necessary to work with irregularly stored data.

In this blog post, we'll be going through how to parse JSON data stored on SnowFlake, in a case whereby the JSON data object happens to have names that are number-like (see the example below):

{% highlight json %}

{
    "expenses": {
        "1001": {
            "amount": 2345,
            "date": "2021-01-01"
        },
        "1002": {
            "amount": 33478,
            "date": "2020-12-21"
        },
    }
}

{% endhighlight %}

In this example, we have the `"expenses"` object, which contains a set of objects, each of which has the "expense type" as the name (i.e. `"1001"`, `"1002"`).

The goal is to query the `amount` fields for each expense type.

(In an ideal situation, maybe the `"expenses"` could be a list of objects instead, and each could have an `"expenseType"` field).

Let's assume that the data above is stored in a SnowFlake table having 2 columns:

{% highlight plaintext %}

| ID | EXPENSE_INFO                                                                                          |
|----|-------------------------------------------------------------------------------------------------------|
| 1  | {"expenses":{"1001":{"amount":2345,"date":"2021-01-01"},"1002":{"amount":33478,"date":"2020-12-21"}}} |

{% endhighlight %}

SnowFlake provides the [PARSE_JSON](https://docs.snowflake.com/en/sql-reference/functions/parse_json.html) function to ease parsing of JSON data. Using that, the query would normally look something like this:

{% highlight sql %}

SELECT
    ID,
    GET_PATH(PARSE_JSON(EXPENSE_INFO), 'expenses.1001.amount') as expense1001Amount,
    GET_PATH(PARSE_JSON(EXPENSE_INFO), 'expenses.1002.amount') as expense1002Amount
FROM
    TABLE;

{% endhighlight %}


However, this query fails to run successfully because SnowFlake tries to access `"1001"` & `"1002"` as indexes instead of JSON object names.

Luckily, SnowFlake allows the ["Bracket Notation"](https://docs.snowflake.com/en/user-guide/querying-semistructured.html#bracket-notation) when querying.

Using this approach, we can update our query as follows:

{% highlight sql %}

SELECT
    ID,
    parsedExpenseInfo['expenses']['1001']['amount'] as expense1001Amount,
    parsedExpenseInfo['expenses']['1002']['amount'] as expense1002Amount
FROM
    (
        SELECT
            ID,
            PARSE_JSON(EXPENSE_INFO) as parsedExpenseInfo
        FROM
            TABLE
    );

{% endhighlight %}

This version of the query gets us the result we need!

Hopefully you don't have to work with JSON data with number-like field names much, but just in case!

Happy Coding, and stay safe!
