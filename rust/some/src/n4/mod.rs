use std::mem::size_of;

use bit_struct::*;
use serde::Deserialize;


bit_struct! {
    pub struct zError(u16) { 
        overheat: u1,
        _: u1,
        highest_floor: u2,
    }
}

bit_struct! {
    pub struct HouseConfig(u8) { 
        lowest_floor: u4,
        highest_floor: u2,
    }
}

pub union House {
    hc: HouseConfig,
    all_hc_bits: u8,
    ff: f32,
}

enum House1 {
    HC(HouseConfig),
}


fn get_union<T: Sized>(bytes: [u8; size_of::<T>()]) -> T {
    unsafe {
        std::mem::transmute::<[u8; size_of::<T>()], T>(bytes)
    }
}

macro_rules! get_union {
    ($t:ty, $) => {
        
    };
}

pub fn exec() {    
    // let hc = HouseConfig::new(
    //     u4!(0b1001), u2!(0b11));

    // println!("{:b}", hc.raw());
    // println!("{}", hc.raw());

    // ***

    let bytes: [u8; std::mem::size_of::<House>()] = [0b10001101_u8, 0b11_u8, 1_u8, 2_u8];

    let h =  unsafe {
        std::mem::transmute::<[u8; std::mem::size_of::<House>()], House>(bytes)
    };

    let hcc = unsafe { h.hc };
    println!("{:b}", hcc.raw());
    println!("{}", hcc.raw());
}
