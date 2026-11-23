/> A witch brews a potion. Her mana is borrowed inside the spell block. </
seal cauldron -> "Moonbrew";
grain mana -> 10;
dram potency -> 1.5;

seal withdrawalFee -> 18;
scroll cardHolder -> "Long Anisa";
grain balance -> 1000;
grain counter -> 0;

reveal cardHolder;
reveal balance;

/> deposit </
{
  /> a new mana that only exists inside this block </
  grain mana -> 3;
  reveal mana;

  /> no new variable here, so this changes the outer potency </
  potency -> potency * 2;
}

/> the block is over: the outer mana is back, and potency stayed changed </
reveal mana;
reveal potency;