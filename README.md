# BEACON

## Table of Contents
- [BEACON](#beacon)
  - [Table of Contents](#table-of-contents)
  - [🗯️ Overview](#️-overview)
  - [🔧 Technologies](#-technologies)
  - [🔠 Key Elements](#-key-elements)
  - [🛠 Setup and Configuration](#-setup-and-configuration)
  - [💼 Data Access Layer (DAL)](#-data-access-layer-dal)
  - [⏰ Logging](#-logging)
  - [📇 Database](#-database)
  - [🖥 Development Environment](#-development-environment)
  - [🎙️ Communication](#️-communication)
  - [👥 Contributors](#-contributors)

## 🗯️ Overview

Our mission is to design and develop a unified and secure smart home ecosystem leveraging the robustness of Golang technology, facilitated by the Raspberry Pi and Zigbee network. We aim to create a seamless, user-friendly interface that provides homeowners with intuitive control, real-time feedback, and a high level of security for all their smart home devices. By implementing a system with robust token-based authentication and meticulous event logging in a MongoDB database, we commit to ensuring the integrity and confidentiality of user data. Our goal is to enhance the quality of life for our users by delivering a state-of-the-art smart home experience that is secure, efficient, and easily manageable. 

## 🔧 Technologies

- 📣 Languages: Golang, JavaScript, C++, CSS/HTML
- 🌐 Web UI Framework: React JS Library
- 📖 Database: Mongo Database
- 🗝️ Authentication: Token-based for both users and web services
- Version Control: Git

## 🔠 Key Elements 

- React-based web application designed to simulate a smart home system, enabling users to monitor and control various devices and users.
- Restful API to handle and authenticate front-end requests as well as facilitate DAL functionality.
- Database Access Layer to interface database and update IoT devices.
- Zigbee Network to facilitate secure communication between Raspberry Pis an associated devices.
- AES BlockChain to ensure data integrity, security, and decentralization of device communication.

## 🛠 Setup and Configuration

- ⚙️ Configuration: 
- ....list which file has the setup

## 💼 Data Access Layer (DAL)
- CRUD (CREATE, READ, UPDATE, DELETE) operations for users and devices within a smart home system
- Update functionality to change IoT device states via message serialization and broadcasting.
- IoT device state changes are logged and retrievable

## ⏰ Logging📝

- All IoT device updates are logged for monitoring, analysis, and auditing purposes.
- Logged data is stored within a database collection for retrieval from the API.

## 📇 Database
- Contains user collection comprised of name, role, and login credentials.
- Stores various device collections which correspond with IoT devices.
- Log collection to track IoT device updates.

## 🖥 Development Environment

- 🛠 IDE: [JetBrains GoLand](https://www.jetbrains.com/go/) 
- ...specify the file that has the instructions to set up the dev env 

## 🛠 Deployment and CI/CD Workflow

### Build Workflows

- The repository is configured with build workflows to compile binary files and store them as zipped artifacts in the GitHub Actions artifact store.
- Locations of the build files:
  - For the Golang backend: [go-build.yaml](https://github.com/PSUBeacon/CMPSC488SP24SecTuesday/blob/main/.github/workflows/go-build.yaml)
  - For the React frontend: [fe-build.yaml](https://github.com/PSUBeacon/CMPSC488SP24SecTuesday/blob/main/.github/workflows/fe-build.yaml)

We utilize cross-platform compilation to compile the backend code for both Raspberry PI 3 and Raspberry PI 4 versions.

#### Build for Raspberry PI 3

- To build the binary file for Raspberry PI 3, we use the following flags with `go build`: `GOOS=linux GOARCH=arm GOARM=7`

#### Build for Raspberry PI 4

- To build the binary file for Raspberry PI 4, we use the following flags with `go build`: `GOOS=linux GOARCH=arm64`

#### Service Deployment

- GitHub Actions is configured to SSH into machines using IP addresses stored in the Secrets storage. It copies artifacts to these end machines using an SSH key, which is also located in the Secrets store.

### Service Testing

#### Test Workflow

- The repository has a test workflow configured to test the software using GitHub Actions.
- Location of the test workflow files:
  - For the Golang backend: [test-api.yaml](https://github.com/PSUBeacon/CMPSC488SP24SecTuesday/blob/main/.github/workflows/test-api.yaml)

#### Testing Details

- GitHub Actions is used to run automated tests of the software.
- The build pipeline performs the following tasks:
  - Spins up a MongoDB server.
  - Imports database definitions into the MongoDB server running within a container on GitHub Actions.
  - Builds and runs binaries in the container.
  - Executes API requests to test various endpoints and overall software functionality.
  - Utilizes the Testify library to facilitate writing and running tests.


## 🎙️ Communication

- 📪 Communication Platform: All team communications will be managed via MS Teams.

## 👥 Contributors

<div align="left">
   <table>
  <tr style="display: table-cell">
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Besjana Kubick
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Nikolay Sizov
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/cbussom">
     <img src="https://github.com/cbussom.png?size=50">
      Cameron Bussom
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Charles Angel Langley
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/devv64bit">
     <img src="https://github.com/devv64Bit.png?size=50">
      Dev Patel
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/dharmik529">
     <img src="https://github.com/dharmik529.png?size=50">
      Dharmik Patel
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/dravyaaa">
     <img src="https://github.com/dravyaaa.png?size=50">
     Dravya Patel
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/EricJ-code">
     <img src="https://github.com/EricJ-code.png?size=50">
      Eric John Estadt
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Lasha Kaliashvili
   </a>
   </td>
  </tr>
  <tr style="display: table-cell">
    <td style="display: block">
    <a href="https://github.com/mabraham2o24">
     <img src="https://github.com/mabraham2o24.png?size=50">
      Mahima Susan Abraham
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Mark Douglas Vernachio
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Mikeil Uglava
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Mohamed Kareem Chikani
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Nicanor Sanderson
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://github.com/github.png?size=50">
      Richard Paul McDowell
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/Taazkir">
     <img src="https://github.com/Taazkir.png?size=50">
      Taazkir Nasir
   </a>
   </td>
    <td style="display: block">
    <a href="https://github.com/TreasureAD">
     <img src="https://github.com/TreasureAD.png?size=50">
      Treasure Davis
   </a>
    </td>
    <td style="display: block">
    <a href="https://github.com/">
     <img src="https://upload.wikimedia.org/wikipedia/commons/5/59/Empty.png?size=50" style="height: 50px">
   </a>
   </td>
  </tr>
  
</table>
</div>

