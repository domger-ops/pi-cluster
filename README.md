### Raspberry Pi 4 Home Lab 
update: jan 12.23

**Objective:**
I am embarking on the creation of a home tftp server using Raspberry Pi 4 with the primary aim of gaining hands-on experience in networking services. The ultimate goal is to establish an accessible K3s cluster capable of running multiple applications encapsulated within Docker images. This project serves as a practical exploration of containerized environments and networking intricacies, enhancing my skills in the realm of distributed systems and cloud-native applications.

**Hardware:**
- Raspberry Pi 4 Model B - 8GB
- Case - Heat Sink - Fan - Power Cable w/ switch
- WiFi dongle

**Key Accomplishments**
- Successfully assembled the Pi Server, installed the Ubuntu Server OS, and established SSH connectivity.
- Configured a basic Apache server housing a privately accessible "hello-world" app.
- Implemented and configured K3s, a lightweight Kubernetes distribution, for efficient container orchestration.
- Developed a Golang web application and conducted thorough testing within the Kubernetes containers.
- Created a Docker Image encapsulating the Golang application for streamlined deployment.
- Deployed Docker Images to the containers, ensuring seamless application hosting.
- Secured accessibility within a private network, bolstering privacy and control.
- Established seamless connectivity to the corresponding GitHub repository, fostering version control and collaborative development.
- Implemented a GitHub Actions workflow to automate the building and versioning of Docker images upon pushing changes to the Git repository.

**Goals**
- CD (eventually)
- Publicly accessible (challenge: security)
- Set up different types of k8s deployments (ex: blue/green, etc)

**TL/DR**
raspberry pi home server for me to learn on and practice with
