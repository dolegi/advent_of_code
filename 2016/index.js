const parseInput = input => {
    const match = input.trim().match(/^([LR])(\d+)$/);
    if (!match) { throw new Error('Invalid input'); }
    return [match[1], parseInt(match[2])];
};

const Traveler = () => {
    let x = 0;
    let y = 0;
    let state;

    var North = () => ({
        turn(d) {
            state = d == 'L' ? West() : East();
        },
        move(n) { y = y + n; }
    });

    var South = () => ({
        turn(d) {
            state = d == 'L' ? East() : West();
        },
        move(n) { y = y - n; }
    });

    var East = () => ({
        turn(d) {
            state = d == 'L' ? North() : South();
        },
        move(n) { x = x + n; }
    });

    var West = () => ({
        turn(d) {
            state = d == 'L' ? South() : North();
        },
        move(n) { x = x - n; }
    });

    state = North();

    return {
        go(input) {
            const parsed = parseInput(input);
            state.turn(parsed[0]);
            state.move(parsed[1]);
        },
        delta: () => Math.abs(x) + Math.abs(y)
    };
};

const solve = input => {
    const traveler = Traveler();
    input.split(',').forEach(traveler.go);
    return traveler.delta();
};

x = solve('R5, L2, L1, R1, R3, R3, L3, R3, R4, L2, R4, L4, R4, R3, L2, L1, L1, R2, R4, R4, L4, R3, L2, R1, L4, R1, R3, L5, L4, L5, R3, L3, L1, L1, R4, R2, R2, L1, L4, R191, R5, L2, R46, R3, L1, R74, L2, R2, R187, R3, R4, R1, L4, L4, L2, R4, L5, R4, R3, L2, L1, R3, R3, R3, R1, R1, L4, R4, R1, R5, R2, R1, R3, L4, L2, L2, R1, L3, R1, R3, L5, L3, R5, R3, R4, L1, R3, R2, R1, R2, L4, L1, L1, R3, L3, R4, L2, L4, L5, L5, L4, R2, R5, L4, R4, L2, R3, L4, L3, L5, R5, L4, L2, R3, R5, R5, L1, L4, R3, L1, R2, L5, L1, R4, L1, R5, R1, L4, L4, L4, R4, R3, L5, R1, L3, R4, R3, L2, L1, R1, R2, R2, R2, L1, L1, L2, L5, L3, L1')

console.log(x)
