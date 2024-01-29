# Golang test for developer (integrations)

## Description

Implement simple integration grpc service.

# Quick Start
```
docker build -t integration_app .  
docker run --rm -p 8080:8080 integration_app
```

## gRPC endpoints:
### GetBalance (payload: `exampleGetBalance.json`) 
-- validating request fields    
-- making request to service method `GetGames`  
-- finding game according to `game_id`  
-- sending responce data

### SendBet (payload: `exampleSendBet.json`)
-- validating request fields  
-- subtracting amount from balance and updating in-memory if player has sufficient balance  
-- returning updated balance  


