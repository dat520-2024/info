# Guide to get Unix Account

## UiS Unix System

Most lab assignments can be performed on your local machine.
If you already run Linux or macOS on your laptop, you should be ready to go.
Linux and macOS are to a large extent relatively similar at the command level.
If you are running Windows, please consult the instructions [here](setup-wsl.md).

However, we will also be using some machines in our Linux lab, named `gorina1` - `gorina3`, through remote access; see below for more information about this.
All necessary software for this course should be installed on these machines.

Some of the assignments includes a networking part, requiring you to run your code on several machines.
This can be conveniently done using the machines mentioned above.
To be able to log into these machines you will need an account on the Unix system.

For more rapid testing, you may also run networking code from different ports on your local machine, i.e. `localhost`.

### Task: Sign up for Unix Account - Do It Now

Some of you may already have done this one:
You will need a Unix account to access machines in the Linux lab.
Get an account for UiS’ Unix system by following the instructions [here](https://user.ux.uis.no).
Be sure to read the section **Using the UNIX system**.

*It may take a while before you get access, but you can continue learning while you wait.*

### Accessing UiS Unix Machines

To access the Unix system from outside UiS campus, you need to use one of the following jump servers:

```log
ssh1.ux.uis.no
ssh2.ux.uis.no
ssh3.ux.uis.no
```

From these jump servers you can access other Linux machines such as `gorina1` - `gorina3`.
