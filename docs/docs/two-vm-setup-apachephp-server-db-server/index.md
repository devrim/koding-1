---
layout: doc
title: Simple two VM setup
permalink: /docs/two-vm-setup-apachephp-server-db-server
parent: /docs/home
---

# {{ page.title }}

Configuring your custom development environment requires building a **Stack Template file**. As you configure your stack file, you can add the number of machines required, set the machine types, and choose what applications/packages to be installed on each VM.

While you can setup all that using the visual interface, understanding your **Stack** template YAML file is a great advantage which will give you more control on the development environment VMs or the **Stack** you are creating.

![two-vm-setup.png][1]

## Example

Let's create a stack setup with 2 AWS Micro VMs to learn more about stack templates:

1. Apache &amp; PHP server
2. MySQL server
Go ahead and start a new stack, and in the installed services screen, check the below options:

### Stack template

Here is the **Stack Template** file, we added commands under the `user_data` section for both VMs to install the required packages:

```yaml
# Here is your stack preview
# You can make advanced changes like modifying your VM,
# installing packages, and running shell commands.

provider:
  aws:
    access_key: '${var.aws_access_key}'
    secret_key: '${var.aws_secret_key}'
resource:
  aws_instance:
    example_1:
      instance_type: t2.micro
      ami: ''
      tags:
        Name: '${var.koding_user_username}-${var.koding_group_slug}'
      user_data: apt-get -y install mysql
    example_2:
      instance_type: t2.micro
      ami: ''
      tags:
        Name: '${var.koding_user_username}-${var.koding_group_slug}'
      user_data: apt-get -y install apache php
```

### Stack file sections

Let's look at the auto generated **Stack Template** file parameters

#### 1. Provider Section

```yaml
# Here is your stack preview
# You can make advanced changes like modifying your VM,
# installing packages, and running shell commands.

provider:                                # Starts with the header name "provider:"
  aws:                                   # Provider source, in our example it's AWS
    access_key: '${var.aws_access_key}'  # Your access key (pulled from the credentials you provided)
    secret_key: '${var.aws_secret_key}'  # Your secret key (pulled from the credentials you provided)
```

#### 2. Resources section

```yaml
resource:                                            # Starts with the header name "resource:"
  aws_instance:                                      # All VMs are AWS Instances
    example_1:                                       # First VM name
      instance_type: t2.micro                        # First VM type
      ami: ''                                        # AWS AMI
      tags:                                          # Tags
        Name: '${var.koding_user_username}-${var.koding_group_slug}'
      user_data: apt-get -y install mysql-server-5.6 # Command to run when creating the instance..
                                                     # ...here we're installing MySQL server
    example_2:                                       # Second VM name
      instance_type: t2.micro                        # Second VM type
      ami: ''                                        # AWS AMI
      tags:                                          # Tags
        Name: '${var.koding_user_username}-${var.koding_group_slug}'
      user_data: apt-get -y install php5 apache2     # Command to run when creating the instance
                                                     # (here we're installing Apache and php)
```

### Modifying the Stack Template file

You can directly edit the **Stack Template** file, below we changed the VM names and added a couple of command lines to run when the **Stack** is first created by our team members.

Check the changes we made to the **Stack Template** file to further customize our **Stack**:

First VM name changed from "_exmaple_1_" to "_backend_db_":

```yaml
resource:
  aws_instance:
    backend_db:     #originally: "exmaple_1"
      instance_type: t2.micro
      ami: ''
      tags:
```

* * *

We are going to run multiple commands, using pipe "&#124;" permits multiline. We run 'apt-get update' first, then install mysql server:

```yaml
user_data: |-
        apt-get update
        apt-get -y install mysql-server-5.6
```

* * *

Second VM name changed from "_exmaple_2_" to "_apache server_"

```yaml
apache_server:     #originally: "exmaple_2"
      instance_type: t2.micro
      ami: ''
      tags:
```

* * *

Again we use pipe "&#124;" to write multiline commands for our second VM "apache_server".

We run '**apt-get update**', install apache &amp; php5 and create a folder called "_custom-folder_" in the user home directory using the variable that holds the user name '_${var.koding_user_username}_'

```yaml
user_data: |-
        apt-get update
        apt-get -y install php5 apache2
        mkdir /home/${var.koding_user_username}/custom-folder
```

* * *

This is how our complete **Stack Template** file looks after our changes:

```yaml
# Two VMs Stack

provider:
  aws:
    access_key: '${var.aws_access_key}'
    secret_key: '${var.aws_secret_key}'
resource:
  aws_instance:
    backend_db:
      instance_type: t2.micro
      ami: ''
      tags:
        Name: '${var.koding_user_username}-${var.koding_group_slug}'
      user_data: |
        apt-get update
        apt-get -y install mysql-server-5.6
    apache_server:
      instance_type: t2.micro
      ami: ''
      tags:
        Name: '${var.koding_user_username}-${var.koding_group_slug}'
      user_data: |
        apt-get update
        apt-get -y install php5 apache2
        mkdir /home/${var.koding_user_username}/custom-folder
```
_**Note:** You can check other arguments that you can use from the menu **Stack reference -&gt; aws_instance**_

#### If we check our apache VM now, we can see that the apache is running and the _custom-folder_ was created successfully. You may also check your VM IP and use your browser to open the default apache served page.

![two-vm-running.png][4]

Happy Koding!

[1]: {{ site.url }}/assets/img/guides/stack-aws/1-two-vms/two-vm-setup.png
[2]: {{ site.url }}/assets/img/guides/stack-aws/1-two-vms/server01.png
[3]: {{ site.url }}/assets/img/guides/stack-aws/1-two-vms/server02.png
[4]: {{ site.url }}/assets/img/guides/stack-aws/1-two-vms/two-vm-running.png
