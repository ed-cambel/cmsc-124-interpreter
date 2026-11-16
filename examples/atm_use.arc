/> Example: ATM Use </

seal withdrawalFee -> 18;
scroll cardHolder -> "Long Anisa"
grain balance -> 1000;
grain counter -> 0;

reveal holder;
reveal balance;

/> deposit </
{
	grain amount-> 500;
	balance -> balance + amount;
	counter -> counter + 1;
	reveal balance;
}

/> withdrawal </
{
	grain amount -> 500;
	grain charged -> amount + withdrawalFee;
	balance -> balance - charged;
	counter -> counter + 1;
	reveal charged;
}

reveal balance;
reveal counter;