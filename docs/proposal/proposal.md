# School of Computing &mdash; Year 4 Project Proposal Form

## SECTION A

|                     |                       |
|---------------------|-----------------------|
|Project Title:       | Odin                  |
|Student 1 Name:      | James McDermott       |
|Student 1 ID:        | 15398841              |
|Student 2 Name:      | Martynas Urbanavicius |
|Student 2 ID:        | 16485084              |
|Project Supervisor:  | Stephen Blott         |


## SECTION B

### Introduction

Odin is to be a programmable and extendible job orchestration system which allows for the scheduling, management and unattended background execution of individual user created tasks on Linux based systems.

Classically, job schedulers are ad-hoc systems that treat it's jobs as code to be executed. They are primarily concerned with the live editing of files to specify the paramters of what is to be executed, and when it is to be executed. 

In contrast, Odin treats it's jobs as code to be managed before and after execution. While Odin cares about what is to be executed and when it will be executed, Odin is equally concerned with the expected behaviour of your job, which is to be described entirely by the user's code.


### Outline

Odin jobs are to be written in programming languages such as Python, Node.js and Golang. These jobs import libraries and packages written to specificially describe Odin jobs, and this is how Odin is able to determine expected behaviour. In this way, Odin is said to be a programmable system which is significantly extensibile in comparison to other job schedulers.

Once your job is running, we want to know it's behaviour and history of execution. This observability can be achieved through a web facing user interface which displays job logs and metrics. All of this will be gathered through the use of Odin libraries and will help infer the internal state of jobs. For users, this means Odin can directly help diagnose where the problems are and get to the root cause of any interruptions.

Odin is to be a distributed system. This is a stretch goal we have set which we hope will increase the reliability of the system against single-machine failures. This could be achieved using a distributed reliable key-value store like [etcd](https://github.com/etcd-io/etcd) or [Consul](https://github.com/hashicorp/consul) . Both etcd and Consul are written in Golang and use the Raft consensus algorithm to manage highly-available replicated logs.

In all Odin is to be a job orchestration system that supports teams and users who implement the DevOps ideaology.

### Background

During out INTRA placements, we both dealt with various internally written job/task schedulers and other software such as Jenkins and orchestration tools such as Docker Swarm. We came form teams who aimed to implement the best practices of DevOps and we quickly learnt that DevOps is something to be practiced in the tools a team uses. 

When coming up with ideas for the final year project, we shared our experiences with such technologies from our INTRA placements, and we began to note the absence of a true open source alternative to job orchestration that wasnt built on top of exisiting orchestration technologies like Docker Swarm, Kubernetes or Apache Mesos.

### Achievements

Odin will provide a management system based around the periodical execution of jobs written by the user. As previously stated, Odin concerned with the conditions before and after the time of execution, so Odin also provides the observability of jobs through centralised logs and metric visualisation.

These jobs will be written in langauges like Python, Golang and Node.js initially, so users will be provided language support in the form of libraries and packages in their language of choice. These libraries provide functionality that will help infer the internal state of jobs and any given time.

Users will be presented the feature of `job chains`, which is a system of linking one jobs execution to the execution of another. This means by default, each job has what is known as a `job trigger` which can be easily used to kick off any given job.


### Justification

Odin is useful as it gives teams a shared platform for jobs that allows individual members to package their code for periodic execution. Simply, users turn their code into scheduled invidiual jobs, or indeed, job chains. Specificially, Odin would be of benefit to teams developing on Linux machines. 

Odin is useful to users and teams from an observability point of view, providing a set of metrics and variable insights which will in turn lend transparency and understanding into the execution of all system ran jobs.


### Programming language(s)

- Golang 
    It is likely that the bulk of Odin is written in Golang. This project will require a language than runs well at scale, and Golang does exactly that. We expect to build the command line tooling, scheduling algorithms, and any API's with Golang. Libraries we are looking at for such things include [cobra](https://github.com/spf13/cobra) and [chi](https://github.com/go-chi/chi) Along with this, Golang supports two features that make concurrency easy out of the box: goroutines and channels. Concurrency is especially useful for building software that can take full performance advantage of multiple cores or process threads, and thus it makes sense for us to pursue this as our main language.

    As mentioned earlier, we intend to write packages to support the writing of Odin jobs. Likely one language we will initially support is Golang.

- Node.js (Typescript / Javascript)

    We intend to use Node.js (in the form of Typescript) to build a web facing user interface which will be visualise job metrics, job logs, and observe variable values during and after job execution. The web application front of Odin is responsible giving users a direct insight into the execution of their jobs, and we feel a framework such as Angular is perfect for building this interface. Along with this, we intend to write Node.js package support for Odin jobs.

- Python3

    We intend to write Python package support for Odin jobs.


### Programming tools / Tech stack

As mentioned earlier, the job scheduler system will be written in Golang and will be compiled with its standard compiler. chi (linked above), will be used for building the centralised job scheduler API and cobra (linked above) will be used to build the command line tooling for Odin.

Explained previously, Odin packages will be written in their corresponding languages i.e Python library will be written in Python.

The web interface of Odin will be made using Node.js as the backend, AngularJS as the frontend framework, MongoDB for storing information and ExpressJS as the web server.

^^ we should explain the purpose for all of these

The possibility of distributing Odin across multiple machines can be achieved through the use of etcd or consul, which both implement the Raft consensus algorithm.

### Hardware

Odin will be designed to run on any core distribution of Linux. The platform on which we will be testing the system is Ubuntu Linux.

### Learning Challenges

While both of us have used Golang, neither of us have used it in a project of this scale. Luckily we do have some experience with the libraries of cobra and chi, so this should make the learning curve a little less steep moving forward. A key area in the metrics end of this proposal is the idea of observability, a measure of how well internal states of a system can be inferred from knowledge of its external outputs. While both of us have some experience in various monitoring stacks, we have never extended our reach into the field of observability. The possibility of implementing Odin as a distributed system poses some of most potentially interesting challenges to the project. Distribution often introduces more particular issues from the perspective of testing and validation.

### Breakdown of work

#### Student 1 - James McDermott

James plans to take the helm on the command line interface for Odin, having more experience designing and building robust command line tooling in Golang and other languages.
James will work on the on the central job scheduler system and Odin API with Martynas, and will also work to build the Odin job executor system which will store values for the front end observability functionality. James will focus on developing the Odin Python library with Martynas and will solely focus on the Golang library.

#### Student 2 - Martynas Urbanavicius

Martynas plans to take the helm on the web front, having more experience with the likes of Angular and ExpressJS. Along with this, Martynas will have the job of implementing the visualisation of the metrics and observability of values given from job execution. Along with this Martynas will work on the central job scheduler system and Odin API with James. Martynas will focus on developing the Odin Python library with James and will solely focus on the Node.js library.
