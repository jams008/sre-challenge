# SRE Challenge: Pixel Pet Virtual Pet Microservice

## Overview

Build a microservice that manages virtual pets with various stats (happiness, hunger, energy). While the core application can be built using AI tools, we're particularly interested in the human aspects of operating and maintaining the service - the critical thinking, decision-making, and communication skills that set great SREs apart.

## What We're Looking For

### Design Thinking
- How do you approach operational challenges?
- What tradeoffs do you consider when making architectural decisions?
- How do you balance reliability with complexity?

### Operational Decision-Making
- What metrics do you choose to monitor and why?
- How do you determine alert thresholds?
- What failure scenarios do you prepare for?
- How do you prioritize reliability improvements?

### Communication & Documentation
- Can you clearly explain your operational decisions?
- How do you document system behavior and failure modes?
- How do you structure runbooks and operational guides?

## Technical Requirements

### 1. Core Service
- Build a basic CRUD API for managing virtual pets
- Use MongoDB for data storage
- Use an LLM to assist with the application code
- Focus on making thoughtful operational decisions

### 2. Operational Requirements
- Containerize the application
- Implement monitoring and logging
  - Choose meaningful metrics
  - Set appropriate log levels
  - Consider what to alert on
- Create runbooks for common scenarios
- Document your decision-making process

### 3. Automation
- Set up CI/CD with GitHub Actions
- Create operational tooling
- Document automation decisions

### 4. Local Development
- Provide Docker-based environment
- Include basic development tools
- Clear setup instructions

## Project Structure

We've provided a basic structure to get you started:

```
.
├── .github/
│   └── workflows/        # GitHub Actions workflows
├── src/
│   ├── models/          # Data models
│   ├── routes/          # API endpoints
│   ├── utils/           # Utilities and helpers
│   └── tests/           # Test suite
├── docker-compose.yml   # Local development setup
├── Dockerfile          # Application container
└── README.md          # Documentation
```

Feel free to modify this structure based on your design decisions, but be sure to explain your choices.

## Bonus Points
- Infrastructure as Code
- Advanced deployment strategies
- Innovative monitoring solutions
- Thoughtful scaling considerations

## Getting Started

1. Fork this repository
2. Use an LLM to help build the core application
3. Focus on operational excellence
4. Document your thought process

## Submission

Your submission should demonstrate:
- Clear reasoning behind operational decisions
- Thoughtful consideration of failure scenarios
- Well-structured documentation
- Effective communication of technical concepts

Remember: While AI can help write code, we're interested in your:
- Critical thinking
- Problem-solving approach
- System design decisions
- Operational insights
- Communication clarity

Good luck! We're excited to see your unique approach to building and operating a reliable service.