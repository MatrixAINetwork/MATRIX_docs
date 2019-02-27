### DOCER-BASED SINGLE-NODE CHAIN STARTUP - Detailed steps

This guide aims to demostrate how to start up a single node chain as well as MATRIX online wallet and blockchain explorer


Step 1: Install and Run Docker with a prompt: Docker Desktop is running

Step 2: Open a cmd window, and enter the following command

Step 3: Pull the mirror images

- docker pull dockermatrix123/matrix:matrix 
- docker pull dockermatrix123/matrix:man

Step 4: Import the chain mirror image - docker load < dockermatrix123/matrix:matrix

Step 5: Import the wallet and explorer mirror image - docker load < dockermatrix123/matrix:man

Step 6: Start the GMAN single-node chain - docker run -itd -p 8568:8568 dockermatrix123/matrix:matrix

Step 7: Exeucte: docker run -itd -p 80:80 -p 3000:3000 -e IP=Your_Ip man:4.0, for example, docker run -itd -p 80:80 -p 3000:3000 -e IP=192.168.3.47 man:4.0

Step 8: Enter the wallet or explorer address in the browser, for example: 192.168.3.47:80 or 192.168.3.47:3000


### DOCER-BASED SINGLE-NODE CHAIN STARTUP - A Quick Method

Step 1: Obtain the file docker-compose.yaml

Step 2: Run docker-compose -f docker-compose.yml up -d to start

Step 3: Enter the wallet or explorer address in the browser: http://127.0.0.1/ or http://127.0.0.1:3000

Note: For the parameter 'IP' in file docker-compose.yaml, it should be changed to 172.0.18.2 for Linux. For windows, 10.0.75.2 shall be remained













