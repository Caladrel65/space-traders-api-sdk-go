This project is to create a wrapper and logic handler for the Space Traders API game, specifically the V2 endpoints.

Tasks for this are as follows:

1. Ensure the user is registered (pre-agent, just requires a key/token be present to be used).
2. Create an agent if one does not exist.
2a. If not automatically accepted during creation, accept the starter contract.
2b. If not already owned during creation, purchase a starting vessel that can accomplish the starter contract (typically mining).
3. Navigate to a good place to mine within the same system.
4. On a loop, mine until full, navigate to a place to sell or turn in for contract, do so, then return to mine again.
4a. If returning doesn't work (no more can be mined), find a new location to mine.
4b. Once enough funds are available, purchase mining drones or other infrastructure to speed up mining.
4c. Once using multiple miners, utilize "freighter" ship roles, where one simply carries mined materials back and forth.
4c1. To do this, have "miner" roles pass cargo off to "freighter" roles.