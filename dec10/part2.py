from scipy.optimize import milp, LinearConstraint, Bounds
import numpy as np

class Machine:
    def __init__(self, nr_lights: int, buttons: list[list[int]], target: list[int]):
        self.nr_lights = nr_lights
        self.buttons = buttons
        self.target = target


def part2(file_name: str) ->int:
    with open(file_name) as f:
        lines = f.readlines()

    sum = 0
    for  line in lines:
        line = line.strip()
        parts = line.split()
        nrLights = len(parts[0])-2
        combinations = []
        for part in parts[1:-1]:
            buttons = [0]*nrLights
            bs = [int(x) for x in part[1:-1].split(",")]
            for b in bs:
                buttons[b] = 1
            combinations.append(buttons)
        target = [int(x) for x in parts[-1][1:-1].split(",")]
        machine = Machine(nrLights, combinations, target)
        smallest_presses = find_fewest_presses(machine)
        sum += smallest_presses
    return sum

def find_fewest_presses(machine: Machine) -> int:
    # minimize the sum of presses
    # Conststraings: machine.buttons x = machine.target, x >= 0
    A_eq = np.array(machine.buttons).T
    b_eq = np.array(machine.target)
    costs = np.ones(len(machine.buttons))

    constraints = LinearConstraint(A_eq, b_eq, b_eq) # Ax = b
    bounds = Bounds(lb=0)  # x >= 0
    integrality = np.ones(len(machine.buttons)) # all variables are integers

    # Use MILP solver, Mixed Integer Linear Programming
    result = milp(costs, constraints=constraints, bounds=bounds, integrality=integrality)
    if result.success:
        print("Optimal value:", result.x)
        nrPresses = int(sum(result.x))
        return nrPresses
    else:
        raise ValueError("No solution found")

if __name__ == "__main__":
    result = part2("input")
    print("Result:", result)