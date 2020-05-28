# Documentation


1. [The Odin Scheduling Format]()
2. [The Odin CLI]()
    - [Getting help]()
    - [Generate required job files]()
    - [Deploying a job]()
    - [Removing a job]()
    - [Show a jobs log]()
    - [Describe running jobs]()
    - [Modify running jobs]()
    - [Link jobs together]()
    - [Unlink jobs from each other]()
    - [Recover job files]()
    - [View job stats]()
    - [Add more Odin Engine nodes]()

3. [The Odin Software Development Kits]()
    - [Python SDK]()
    - [Go SDK]()
5. [The Odin Observability Dashboard]()
6. [The Odin Engine as a Distributed System]()

---

<br/>

## 1. The Odin Scheduling Format

The Odin Scheduling Format is heavily reliant on the simplicity of the English language. 

Rather than fumble around with the cron schedule string syntax, we have created a more human readable scheduling format that leverages a subset of specific keywords to form a robust schedule string syntax.

Here are some examples of acceptable Odin schedule strings:

![](https://lh3.googleusercontent.com/NLOoG7Pw_3aqcf4nK_UU4334lMBzmJtQUnvbkmYbg1R_njOClhjCui-ZMF33pQA9VevZqFon7gTr2oHj0TORRaQ7PKIDcG4bxvCs14fE-EjRabHnGSk6nR-N9EBeZYgdsa6NyDON)

<br/>

## 2. The Odin CLI



### Getting help
To view all available Odin CLI commands you can simply run:

```odin --help```

To view available flags and recommended usage of a specific command, run:

```odin <command_name> --help```

<br/>

### Generate required job files
Every Odin job is defined by two separate files:

- The user code (a `.py`, `.go` or `.js` file) used to define a jobs operation(s)
- Supplementary YAML (a `.yml` or `.yaml` file) which contains associated job metadata

We can generate the appropriate files for our job like so:

```odin generate -f amzn_stocks.yml -l python```

`-f` specifies a new YAML file for the job in the format: ```<jobName>.yml``` The value passed here will also be used to name the user code in the form of:
- ```<jobName>.py```
- ```<jobName>.go```
- ```<jobName>.js```

`-l` defines the language used for the job. The acceptable values for the this parameter are:
- python
- go
- node

In the first example we generated two files:
- ```amzn_stocks.yml```
- ```amzn_stocks.py```

```odin generate -f backup.yml -l go```

Similarly in the above second example we generate two files again, but this time the extension of the user code file is different:
- backup.yml
- backup.go


The contents of the YAML file are purposely bare in parts:

![](https://lh3.googleusercontent.com/WXgZUYX3X2FeT6hiIYtm9P-6EhkjqdahAKXDOjpOMJLcnXbeSsMb2pV3961r4nU3s3tIiuSpS6AoD0tG9_J_0w2av1QS1UCRzbjekzNXdqD6YAQQ0Nri4JpYRz87U_9s1Td5ylcu)

<br/>

We leave the completion of the name, description and schedule fields to the user. The name and descriptions can be set at the user’s discretion, and we advise you to consult the Odin scheduling format in section 4.1 of this document when setting the schedule string in the YAML file.

The generated user code file is purposely left completely empty for the user to fill their code in.

<br/>

### Deploying a job

Once you have generated your job files and filled them out as appropriate, deploying your job is simple. Here’s a look at the files used during this example deployment.

**YAML: amzn_stocks.yml**

![](https://lh6.googleusercontent.com/FtapSI-cD4PJ-jxR3Hq3y3_1vNDxNJW6Fs_KIoueYVIFTOwW0y4lv1Dp664GH1RpcIQqOdj7RqV3cVuFAbWhNjM-sUfFppvI611HzFOUB3tX6NOHXPZg_itLh0fDrQuLFfTFafr3)


**USER CODE: amzn_stocks.py**

![](https://lh5.googleusercontent.com/5cIlXax9oJ2-RGfzxb1uNYXkWsi6s4QsQdfAPidipDbuUnQpts83hduh3F6p7Zxz_V1uPYt_hQqaLkfJ9mVzqTgjqF8m0QUmqITy5f6O9UMAR9jOga6bIoNoST-e4-A6Igry8twp)


We simply deploy the above job with the following command:

```odin deploy -f amzn_stocks.yml```

Where `-f` specifies the YAML file for the job.

We can check if this deployment was successful by running:


```odin list```

Running this command gives us the following output:

![](https://lh4.googleusercontent.com/FHVsMnF3Xgl9ab_vk8P0US2v0m6hD557jVz-OtC17UHtBSPGl_jI_FZPd2O6vjsfsw67WS268Qauea4Vq7sjuOxgPNTjmMR1PUKfUwWlYpaY3PnMfg8iupKM16jgVRNo0DOOB1P_)

We can see that we have successfully deployed our job! The `deploy` command is aliased to both `dep` and `add`. The `list` command is aliased to `ls`.

<br/>

### Removing a job

We can remove any job by fetching it’s ID displayed with `odin list` and using using it in the following command:


```odin remove -i a37b187f5930```

We can check if this removal was successful by running `odin list` again:

![](https://lh3.googleusercontent.com/H0KxjytMeWtltHxD1ZQLETotQhQSytNrC72NbyyFjW6a8UkjKBHi0DPBSzkONyVcz5gKR-iAaTaNFg2Ic1pwzfNtzHH6mwlOpekSM3j9lQVHRwnAr-OyRxfaZ38Pt43U4ic54uVm)

We can see that we have successfully removed that job! The `remove` command is aliased to `rm`.

<br/>

### Show a jobs log

We can view the execution logs of any job by fetching it’s ID displayed with `odin list` and using using it in the following command:

```odin log -i a37b187f5930```

This will yield an output which looks like this:

![](https://lh6.googleusercontent.com/f5TWZLYNya7QcisxIShmeYC_ERh1UNW-0yxVNMd4bgDAEN1sGp-L6CvY8CWZTN-JDm5mL-neO4M6AEPJDFxNr5fE2rnv21YK4Ngy7OTL7yRWbahWOkvIFwH92y_CClpStZA2_w9S)

In the above image we can see:
- the time of execution
- the message level (info or warning)
- the message attached to this log (exec or failed)
- the user code of the job that was executed
- the uid and gid of the user who deployed that job
- the language of the job executed
- the node on which it the job was executed

<br/>

### Describe running jobs

We can view the execution logs of any job by fetching it’s ID displayed with `odin list` and using using it in the following command:

```odin describe -i a37b187f5930```

This will yield an output which looks like this:

![](https://lh6.googleusercontent.com/N4rtQr4BZdrYo5ORZXG8B58m2WFF7ObeNJdYlYLEj2cSdBG9XQDkJIsszKjBp3oVog3bijWPpYAZvdJPZ-Bsqft5jUF0PDdzq0SlaIj15BnlODBy5YStm0Z7C70TTqeEjQSqudSw)

We can see that we have successfully viewed that job description! The `describe` command is aliased to `desc`.

<br/>

### Modify running jobs

We can modify a jobs name, description and schedule with the odin modify command. Once more we must fetch the ID to specify the job we wish to modify. Once you have the ID you can change the name, description and schedule with the `-n`, `-d` and `-s` flags respectively. Below are some examples of the odin modify command.

Modify the description alone:
```odin modify -i a37b187f5930 -d "check my stocks"```

Modify the schedule alone:
```odin modify -i a37b187f5930 -s "every hour"```

Modify the name alone:
```odin modify -i a37b187f5930 -n "stock_check"```

Modify the description and the schedule:
```odin modify -i a37b187f5930 -d "a job to check my stocks" -s “everyday at 18:00”```

As long as you remember to include quotation marks around the new value for any of the values you want to change, the modify command should work and you can view your successful changes with the odin list command once more. The `modify` command is aliased to `mod`.

<br/>

### Link jobs together

When you run the `odin list` command you will see a LINKS heading, which is empty by default. Links in Odin reference a concept where the successful execution of one job will in turn call another job. 

Let’s say the `amzn_stock` job runs every hour and stores the information in a file in the user home directory. If this job executes successfully we would like it to call another job, `email_file`. 

The schedule `never` is totally acceptable in the Odin Engine - it essentially denotes that a job will never execute by itself but will only execute if linked to another job.


Once we deploy this job running `odin list` will give us this output:

![](https://lh5.googleusercontent.com/T2ThsX46V50z_kCSG2fPjBfHfYBfJ0cfd1KmT2Fb-Hp_fYyFxBMGJXP2pX9m27iFRVfD5U-AxKphQBCDGlXbzCkEU69Ln9gGwDTA7_Upnm0SyYkTOgiT9sWQIAb74YDN1kWF6Yqu)

We can link the stock check job to the email job with the following command:

```odin link -f a37b187f5930 -t f0aedde327f3```

Running odin list will show us that the link between these jobs now exists:

![](https://lh6.googleusercontent.com/i0Ci_3_YDcWMFe-RLF2XbN2N4va5CxfnEW_LM3MBg-YNbVqbMNibIBYIbEGv9eDStr3BU1zlolVrYYW4b2NvO8Ego6P9X2trsYl35dnxiICctpkRuu1sSe73LXdCfWh4Km63ESoS)

<br/>

### Unlink jobs from each other

We can undo the work done by the `odin link` command using the unlink command like so:

```odin unlink -f a37b187f5930 -t f0aedde327f3```

This decouples the two jobs into independent tasks once more.

<br/>

### Recover job files

While working, let’s say you accidentally removed the user code and/or the YAML file. There’s no need to fret, the Odin Engine has already stored them for you in case this happens, and it will continue to execute them for you. 

If you need to recover these files for any reason you can fetch the ID by using odin list and use it in the command below:

```odin recover -i a37b187f5930```

Running `ls` in your current directory will show that the files have been fully restored.

<br/>

### View job stats

You can directly view statistics with the `odin stats` command by specifying the ID of the job you wish to view. In the case of the earlier `amzn_stock.yml` job we can type:

```odin stats -i 6d7341fa3fea```

This gives us an an output which will look something like this:

![](https://lh5.googleusercontent.com/9ghEtYC48WPclmqqnRCMYPqlV1TptfC3r4MPL4GnD4DkgO_oi1R4mNnx0lyM107lbBl75nRHrG4uWZ0SfBivCM4R0UFCd2CJ6Fv1dreG2tXsohdSW_N2Nu3cpV_rVz6OcnMvkKnk)

These values from the code are generally tracked thanks to the Odin Software Development Kits. You can learn more about the SDK's [here]().

<br/>

### Add more Odin Engine nodes

The Odin Engine leverages the Raft consensus protocol to run as a distributed system. Distributed systems offer reliability to systems in case of a failure, and Odin proves to be an easily scalable and flexible system.

By default, the initial Odin Engine node is called “master-node”. If you want to add nodes to the cluster you can do so simply with the command:

```odin nodes add -n worker1 -a :39391 -r :12001```

Breaking down this command we see that:
- The name (-n) of the new worker node is `worker1`
- The http address port provided for the new worker is `:39391`
- The raft port provided for the new worker is `:12001`

We can add another node like so:

```odin nodes add -n worker2 -a :39392 -r :12002```

We can verify the addition of this new nodes with the following command:

```odin nodes get```

Which returns a list of nodes in the cluster:

![](https://lh3.googleusercontent.com/gdWk9JMp5OqPtSh2ECcWKTkb0knkDa82nMFkCRuNPK3F-oW-1AzNOWUImwOAvU8vAlrX3UowlUN_XEoFKERlq3wm4HMULYx9mNCGr0FsA0lSwmQpTWiuVad0w7MjIWJAIjKUE7UN)

In regards to distributed systems, in particular with raft based systems, it’s advisable to run 3-node clusters or 5-node clusters. This is recommended as:
- 3-node clusters can tolerate a failure in 1 node
- 5-node clusters can tolerate a failure in any 2 nodes

Please consult [here]() in regards to further details on running Odin as a distributed system.

<br/>

## 3. The Odin Software Development Kits

In each of the following subsections we will demonstrate how a language specific SDK is used in conjunction with that language. Operationally each section is the exact same.

<br/>

### Python SDK

The Python SDK can be imported as a pip package using the following command:

```pip3 install pyodin```

You can create the odin object in your python user code like so:


```python
import pyodin
odin = pyodin.Odin(config="your_config_here.yml")
```

It’s important you specify the name of your YAML configuration file here as the Python SDK utilises metadata from it.

From here it’s quite simple, you have three distinctive operations:
- Watch
- Condition
- Result

Let’s look at the `watch` operation with respect to the following code snippet:

![](https://lh4.googleusercontent.com/KbLflWQcbjDiPbFLTy-Yh0cjI7aJxl-OKy2pqOK5mctPFZZIr4vgMxWLg2hY0ZHmKF9sW443aOjbYw_dwun4xDbvC7udJiRtlD_gl1G0mVgfCVC9bjOipYk6dheO_TtGDUTsx148)

This snippet fetches a random url from wikipedia on each execution. The watch operation stores this url along with a string “url fetched” to annotate the variable being stored.

This allows us to better debug our jobs for when they don’t work as intended. If this job failed, we could immediately diagnose whether or not it was because a url wasn’t generated from line 6.

Let’s now look at this code again, but with the `condition` operation added:

![](https://lh3.googleusercontent.com/zYKDEbd-1g3RuloVauQ31Npd2k5iI_VdC53cMGUtD4-E4ubte9Adh_5A5W3dpFUPb6ZSL2ZGjffXsaYs8g6vJKk_aado9kQh8UvlgJhLkXPRcpcwY2oJCWeBzBLPnUB-CDjdGgwu)

This time, we are introducing a new line at line 8. This `condition` operation will store the boolean equivalent to the statement `response.url != ""` and will be annotated by the string “check if url is empty”.

If the statement `response.url != ""` is true then the condition operation will return true and progress to the `watch` operation on line 9.

Finally, we take a look at the `result` operation: 

![](https://lh6.googleusercontent.com/dSG7OPOOZv9PugNaki17S1KBQoJZOeHbS9qrOcNr0I0yDunei1Q-ByOeUkF4FMdq6uuXoFSmaQ7T_kfj1UXIReQmXYRJ_RTttV7g2nvSVFmBquM7Ne-KiZy2AMJJfaw_kwwX0-a-)

This operation acts like a return statement, once it’s executed the job is considered to be over. A successful attempt is denoted by a 0 while a failed attempt is denoted by a 1.

In either case, the result is annotated by a string once more. Once a result operation is run, the code will finish execution.

<br/>

### Go SDK

The Go SDK can be imported using the following command:

```go get github.com/theycallmemac/odin/odin-libraries/go/odinlib```

You can create the odin entry point in your go user code like so:

```go
package main

import (
    "github.com/theycallmemac/odin/odin-libraries/go/odinlib"
)

func main() {
    odin, _ := odinlib.Setup("your_config_here.yml")
}

```

It’s important you specify the name of your YAML configuration file here as the Go SDK utilises metadata from it.

From here it’s quite simple, you have three distinctive operations:
- Watch
- Condition
- Result


Let’s look at the `Watch` operation with respect to the following code snippet:

![](https://lh4.googleusercontent.com/ITRAIkahGTmREUKxsSU7Ig0TzeMsD0clgMTIJFjf-5PaZlh-qDQpLcfCb3ECboUdBpZVdWu4GUkAVmlTGMkgWvwsEPVH2nOFPk4gx8FhL2sbEzQ3Lmv7tMO5jZtTd8uHtdOQyQ2y)


This snippet fetches a random url from wikipedia on each execution. 

The watch operation stores this url along with a string “url fetched” to annotate the variable being stored.

This allows us to better debug our jobs for when they don’t work as intended. If this job failed, we could immediately diagnose whether or not it was because a url wasn’t generated from line 6.

Let’s now look at this code again, but with the `Condition` operation added:

![](https://lh4.googleusercontent.com/pW2L8dlDgI_dm04010raJTCm9hIysZK0SaUwSs8GVTDlL_W6VbuSFOcrlVNSisHxJerxwugCZG6Nt8hJdLe0BuntPZ21xio6oHwX1-86)


This time, we are introducing a new line at line 8. This `Condition` operation will store the boolean equivalent to the statement `url != ""` and will be annotated by the string “check if url is empty”.

If the statement `url != ""` is true then the `Condition` operation will return true and progress to the `Watch` operation on line 9.

Finally, we take a look at the `Result` operation:

![](https://lh6.googleusercontent.com/bgPDnYJc10H7tUXIpRzxqywTrY55ClZ_a8YXRMDlMS7_f0tODoXRfgN7FawDREUd3nQ8OqVym1MFE3bNVW-xpTdvyAKcACTF2unP50cd)

This operation acts like a return statement, once it’s executed the job is considered to be over. A successful attempt is denoted by a 0 while a failed attempt is denoted by a 1.

In either case, the result is annotated by a string once more. Once a result operation is run, the code will finish execution.

<br/>

## 4. The Odin Observability Dashboard

Once running, open the Odin Observability Dashboard in browser. If the configuration is unchanged this can be done by going to `http://localhost:4200`.


![](https://lh6.googleusercontent.com/3tiPEzwQAUQo7xajTQ14bkKi8SfFk2W1NOncO8-wcc6wWo9i6gJFRM-usGWV42uDBxS7Ud0E7pfXt5-wPYUg2M0lq7zAH_JHVZ_Q0dk)

You will be presented with the login page as above. Click the `Log In With Google` buttong and follow the Google Sign In process. If this is your first time logging in, an Odin user account will be made for you automatically.


![](https://lh6.googleusercontent.com/vGEnZyV8metySmnjPF1WPGj0OSlj8ZGeGhqpVfUYen5UakFMB2WTqpNFfAVd3XpWHzzjANqllbh_VdrgoDQjgUbIuuVXoTgGQ7i82nylFcAFzd4LPpxI3MLtBA2cy-i29nZIKoe_)

Once logged in you will be presented with the Odin Observability Dashboard - as shown above.  A list of currently running jobs is shown under the `Jobs List` card. This list acts as a menu where you can select a job for which you want to see observability metrics. Upon selecting a job from the list, the dashboard will update with the job's info and metrics.

After a job is run, the currently selected job metrics will update automatically, as per the difference in metrics cards in between the image above and the image below.

![](https://lh3.googleusercontent.com/Y_fvWZpzBzafg8793ZGcdbOohaXqBCMi__U_X7nncXuTYqq9cUOcyz0Mj-K8jsAVx4PHHXBSPuEoebhfa0iK9LVzrVhRBc4fkLb3ezCYmV4ROPRdK_DeReTv-UAptkqZ1mR6bDYN)

<br/>

## 5. The Odin Engine as a Distributed System

As previously mentioned Odin Engine leverages the Raft consensus protocol to provide a highly available replicated log to run as a distributed system. Distributed systems offer reliability to systems in case of a failure, and Odin proves to be an easily scalable and flexible system. 

Please consult [this section]() in regards to adding nodes to the Odin cluster.

With Raft based systems, it’s advisable to run 3-node clusters or 5-node clusters. This is recommended as:
- 3-node clusters can tolerate a failure in 1 node
- 5-node clusters can tolerate a failure in any 2 nodes


![](https://lh3.googleusercontent.com/gdWk9JMp5OqPtSh2ECcWKTkb0knkDa82nMFkCRuNPK3F-oW-1AzNOWUImwOAvU8vAlrX3UowlUN_XEoFKERlq3wm4HMULYx9mNCGr0FsA0lSwmQpTWiuVad0w7MjIWJAIjKUE7UN)

In the example set out [above](), if the master node fails, a new master node must be elected. Given the set up in the image, the new master node will be worker1 or worker2. This capacity for failure is one of the greatest benefits with distributed systems. At this point however, you may find that Odin CLI operations will seemingly fail as they are made to interact directly with the master node. 

We have built in a --port flag for each Odin CLI operation. This allows users to interface with a node of their choice in the case the master node goes offline. As Raft maintains a replicated log between nodes specifying the node to execute upon will not make any difference outside of having to add an additional flag to your commands.

---
