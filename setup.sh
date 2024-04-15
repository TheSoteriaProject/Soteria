#!/bin/bash

# 1. Git Clone
git clone https://github.com/TheSoteriaProject/Soteria.git
if [ $? -ne 0 ]; then
    echo "Failed to clone the repository."
    exit 1
fi

# 2. Go Build
cd Soteria/Soteria
go build
if [ $? -ne 0 ]; then
    echo "Failed to build the project."
    exit 1
fi

# 3. Add file to /usr/local/bin
sudo cp ./Soteria /usr/local/bin/
if [ $? -ne 0 ]; then
    echo "Failed to copy the executable to /usr/local/bin. Please ensure you have sudo privileges."
    exit 1
fi

# 4. Set permissions
sudo chmod +x /usr/local/bin/Soteria
if [ $? -ne 0 ]; then
    echo "Failed to set executable permissions on /usr/local/bin/Soteria."
    exit 1
fi

echo "Setup completed successfully. You can now run Soteria."
