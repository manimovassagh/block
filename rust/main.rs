mod models;
use models::Block;
use serde::{Serialize, Deserialize};
use std::time::{SystemTime, UNIX_EPOCH};
use std::fs::OpenOptions;
use std::io::prelude::*;
use colored::*;

const DIFFICULTY: usize = 5;

#[derive(Serialize, Deserialize)]
struct Blockchain {
    chain: Vec<Block>,
    difficulty: usize,
}

impl Blockchain {
    fn new() -> Self {
        Blockchain {
            chain: vec![Block::genesis()],
            difficulty: DIFFICULTY,
        }
    }

    fn get_latest_block(&self) -> &Block {
        self.chain.last().unwrap()
    }

    fn add_block(&mut self, mut new_block: Block) {
        new_block.previous_hash = self.get_latest_block().hash.clone();
        println!("{}", format!("Adding new block with index {} and previous hash {}", new_block.index, new_block.previous_hash).cyan());
        new_block.mine_block(self.difficulty);
        self.chain.push(new_block);
        println!("{}", format!("Block added with hash {}", self.chain.last().unwrap().hash).green());
        self.print_blockchain();
    }

    fn is_chain_valid(&self) -> bool {
        for i in 1..self.chain.len() {
            let current_block = &self.chain[i];
            let previous_block = &self.chain[i - 1];

            if current_block.hash != current_block.calculate_hash() {
                return false;
            }

            if current_block.previous_hash != previous_block.hash {
                return false;
            }
        }
        true
    }

    fn print_blockchain(&self) {
        let blockchain_json = serde_json::to_string_pretty(&self).unwrap();
        println!("{}", blockchain_json.yellow());
        let mut file = OpenOptions::new().append(true).create(true).open("blockchain.log").unwrap();
        writeln!(file, "{}", blockchain_json).unwrap();
    }
}

fn main() {
    let mut blockchain = Blockchain::new();

    let block = Block {
        index: 1,
        timestamp: SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_secs(),
        data: "Block 1 Data".to_string(),
        previous_hash: "".to_string(),
        hash: "".to_string(),
        nonce: 0,
    };

    println!("{}", "Starting mining process...".blue());
    blockchain.add_block(block);
    println!("{}", format!("Blockchain is valid: {}", blockchain.is_chain_valid()).magenta());
    blockchain.print_blockchain();
}
