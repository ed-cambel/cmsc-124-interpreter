grain a -> 1; 
{
    grain b -> 2;
    {
        a -> a + b;
        reveal a;           /> expect: 3 </
    }
    reveal b;               /> expect: 2 </
}
reveal a;                   /> expect: 3 </