use std::sync::{Arc, Mutex};

use chrono::{DateTime, Utc};
use common::packets::BanchoPacket;
use dashmap::DashMap;
use tokio::{net::TcpStream, sync::mpsc::{Sender, Receiver}};

use crate::{clients::{self, waffle_client::WaffleClient}, osu::OsuClient};

pub struct ClientInformation {
    pub version: i32,
    pub client_hash: String,
    pub allow_city: bool,
    pub mac_address_hash: String,
    pub timezone: i32
}

pub struct OsuClient2011 {
    pub connection: TcpStream,
    pub continue_running: bool,

    pub logon_time: DateTime<Utc>,

    pub last_receive: DateTime<Utc>,
    pub last_ping: DateTime<Utc>,

    // joinedChannels: 
    pub away_message: String,

    pub spectators: DashMap<u64, Arc<dyn WaffleClient>>,
    pub spectatingClient: Option<Arc<dyn WaffleClient>>,
    
    pub packetQueueSend: Arc<Sender<BanchoPacket>>, 
    pub packetQueueRecv: Arc<Receiver<BanchoPacket>>
}

impl WaffleClient for OsuClient2011 {
    fn get_user(&self) -> common::db::User {
        todo!()
    }
}



impl OsuClient2011 {
    // pub fn as_waffle_client(&self) -> Box<dyn WaffleClient + Send + Sync> {
    //     Box::new(self)
    // }
}