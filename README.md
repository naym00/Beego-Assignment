<h1 align="center">Beego-Assignment - Golang(Beego)</h1>
<p>
  <img alt="Version" src="https://img.shields.io/badge/version-1.0.0-blue.svg?cacheSeconds=2592000" />
</p>
You can get images of cat from this project. There are 4 dropdown buttons, you can get images of cat by selecting or combination of those buttons.
By default it will display random images and fields will be filled like this Order: "Random", Type: "All", Category: "None", Breed: "None".
You can get more images by clicking more+ button.

## Languages used
- Golang

## Framework used
- Beego

## How I completed this project

- First of all, I started learning the core-concepts and syntax of Golang from https://www.golangprograms.com/go-language.html and practised Golang on https://go.dev/play/
- After completing the basics, I switched to Beego and read their official doucumentation from https://beego.vip/
- Finally, I read The CAT API documentaion from https://docs.thecatapi.com/ to understand the API responses and started working on the project.

## Prerequisite

- Install Golang
- Install Beego

## How to Install Golang

#### step 1

- Open your terminal and run the following command to make sure that you’re in the root directory.

```
cd ~
```

#### step 2

- Then use curl to retrieve the tarball.

```
curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
```

#### step 3

- To verify the integrity of the file you downloaded, run the `sha256sum` command and pass it to the filename as an argument

```
sha256sum go1.16.7.linux-amd64.tar.gz
```

#### step 4

- Next, use `tar` to extract the tarball. This command includes the `-C` flag which instructs tar to change to the given directory before performing any other operations

```
sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
```

- In this step, you will set paths in your environment.

```
sudo nano ~/.profile
```

#### step 5

- Then, add the following information to the end of your file:

```
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/home/yourpcname/go/bin
export GOPATH="/home/yourpcname/go/"
export GOBIN="/home/yourpcname/go/bin"
```

- After you’ve added this information to your profile, save and close the file. If you used nano, do so by pressing CTRL+X, then Y, and then ENTER.

#### step 6

- After, check if you can execute go commands by running go version:

```
 go version
```

## How to Install Beego and Create Project

#### step 1

- Download and install beego By this command

```
go get -u github.com/beego/beego/v2
go get -u github.com/beego/beego/v2
```

#### step 2

- Create Beego MVC project by this command

```
bee new project_name
```

- Create Beego API project by this command

```
bee api project_name
```

#### step 3

- Init module

```
go mod tidy
```

## How to run this project

- First of all, clone the project by running the following command

```
git clone https://github.com/naym00/Beego-Assignment.git
```
- Then, open the cloned folder using `VS Code` and run the following command

```
bee run
```
- Finally, open your browser and enter the following URL to check it out

```
http://localhost:8080/
```

## Feedback

If you have any feedback, please reach out to me at nymur@w3engineers.com

## 🚀 About Me

My name is Md. Nymur Rahman and I'm an Intern Softwere Engineer of W3engineers Ltd....
