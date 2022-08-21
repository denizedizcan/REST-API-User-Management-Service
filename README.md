# **REST-API-User-Management-Service**
*This repository contains the necessary codes &amp; info about REST API which is used to manage a user management service. It has some basic CRUD operations on user system. Using Postgres*

# About the project

*The functions of this service are as follows;*

1. List Users
    - API should be able to list all users.
2. Add User
    - API can add users.
3. Show One User
    - API can show one user by id.
4. Delete User
    - API can remove a user by id.
5. Update User
    - Users can update their data

### Prerequisites and Installation

- #### **Docker:** 
  You can install Docker Desktop by following the instructions on the [Docker Desktop website](https://desktop.docker.com/).

  If you are using **MacOS**, you can install Docker by following commands:
    - *`brew install docker`*
    - *`docker run hello-world`*
  
 ---

- #### **PostgreSQL:**

  Plese refer [**here**](https://github.com/denizedizcan/REST-API-User-Management-Service/blob/main/db/README.md) for the initialization scripts.

---

- #### **Go:**
  
  If you haven't done already, You need to install Go by following the instructions on the [Go website](https://golang.org/doc/install).

---
- #### **Project:**

    You have to clone the project from the [Github repository](https://github.com/denizedizcan/REST-API-User-Management-Service/) and run the below script to build the project

    *`git clone https://github.com/denizedizcan/REST-API-User-Management-Service.git`*

---

- #### **Run the Project:**

    You have to make sure that you are in the **same directory as the project**. Then you can run the below script:

    *`docker build -t {your-desired-image-name} .`*

    *`docker run -d --env-file=.env --name {your-desired- container-name} -p {your-desired-port}:8080 {your-desired-image-name}`*

    **Note:** Be careful with the port number. You **can't** run this project as a docker container while **DB_HOST** is set to **127.0.0.1** in .env file

    **Note:** If you want to run the project as a docker container, you have to change the **DB_HOST** to the **IP address of the container** mentioned in the **.env** file. Refer [**here**](https://github.com/denizedizcan/REST-API-User-Management-Service/blob/main/.env) for the .env file.

    **You can get the IP address of the container by running the below command by replecing *{your-container-name}* with a container name you want to inspect**

    *`docker inspect --format '{{ .NetworkSettings.IPAddress }}' {your-container-name}`* 
        
    **Example For docker inspect:**
    *`docker inspect --format '{{ .NetworkSettings.IPAddress }}' postgres-cnt`*


    **Example for running docker container:**

    *`docker build -t project-rllc-img:1.0 .`*

    *`docker run -d --env-file=.env --name project-rllc-cnt -p 8080:8080 project-rllc-img:1.0`*


    Or you can run the project directly by running the below script from the same directory as the project:

    `go run main.go`
