```toml
# post title, required
title = "Welcome to PuGo"

# post slug, use to build permalink and url, required
slug = "welcome"

# post description, show in header meta
desc = "welcome to try PuGo static site generator"

# post created time, support
# 2015-11-28, 2015-11-28 12:28, 2015-11-28 12:28:38
date = "2017-01-01 12:20:20"

# post updated time, optional
# if null, use created time
update_date = "2017-01-01 20:50"

# author identifier, reference to meta [[author]], required
author = "pugo"

# thumbnails to the post
thumb = ""

# tags, optional
tags = ["pugo"]

# draft status, if true, not show in public
draft = false
```

When you read the post, `PuGo` is running successfully.

This post is generated from file `post/welcome.md`. You can learn it and try to write your own article with following guide.

#### Front-Meta

Post's front-meta, including title, author etc, are created by first code section with block **\`\`\`toml ..... \`\`\`**:

```toml
# post title, required
title = "Welcome to PuGo"

# post slug, use to build permalink and url, required
slug = "welcome"

# post description, show in header meta
desc = "welcome to try pugo static site generator"

# post created time, support
# 2017-01-01, 2017-01-01 12:30, 2017-01-01 12:30:40
date = "2017-01-01 12:30"

# post updated time, optional
# if null, use created time
date = "2017-01-01 12:40"

# author identifier, reference to meta [[author]], required
author = "pugo"

# tags, optional
tags = ["pugo"]

# draft status
# if true, it will not be compiled to static page
draft = false
```

#### Content

The content is data after first block. All words will be parsed as markdown content.

```markdown

When you read the post, `PuGo` is running successfully.

This post is generated from file `post/welcome.md`. 

You can learn it and try to write your own article with following guide.

![golang](/media/golang.png)

...... (markdown content)

Markdown is a lightweight markup language with plain text formatting syntax designed
so that it can be converted to HTML and many other formats using a tool by the same name.
Markdown is often used to format readme files, for writing messages in online discussion forums,
and to create rich text using a plain text editor.

```

![golang](/media/golang.png)

Markdown is a lightweight markup language with plain text formatting syntax designed
so that it can be converted to HTML and many other formats using a tool by the same name.
Markdown is often used to format readme files, for writing messages in online discussion forums,
and to create rich text using a plain text editor.
