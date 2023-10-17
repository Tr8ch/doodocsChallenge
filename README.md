# doodocsChallenge
This is the solution to the technical assignment for an internship in Doodocs.
## Getting Started

### Prerequisites

- Go (Golang) installed on your system.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Tr8ch/doodocsChallenge.git

   cd doodocs-backend-challenge

   make run
   ```
## Configuration
The application uses a configuration file to set up various parameters. You can find the configuration file at config.yaml. Make sure to adjust it according to your environment.

## Usage
### API Endpoints
- POST /api/archive/information: Endpoint for archiving information.
- POST /api/archive/files: Endpoint for archiving files.
- POST /api/mail/file: Endpoint for sending email with a file attachment.
### Logging
The application uses a logging library for recording various events. Debug messages are enabled by default.
