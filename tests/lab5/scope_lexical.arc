scroll greeting -> "Hello";

spell greet() {
    reveal greeting;
}

spell rude() {
    scroll greeting -> "Go away";
    greet();
}

rude(); /> expect: Hello </
/> dynamic scope would print Go away </