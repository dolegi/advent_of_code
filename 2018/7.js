let input = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

input = `Step W must be finished before step B can begin.
Step G must be finished before step T can begin.
Step B must be finished before step P can begin.
Step R must be finished before step M can begin.
Step K must be finished before step Q can begin.
Step Z must be finished before step X can begin.
Step V must be finished before step S can begin.
Step D must be finished before step U can begin.
Step Y must be finished before step J can begin.
Step A must be finished before step C can begin.
Step M must be finished before step U can begin.
Step E must be finished before step X can begin.
Step T must be finished before step F can begin.
Step U must be finished before step C can begin.
Step C must be finished before step Q can begin.
Step S must be finished before step N can begin.
Step X must be finished before step H can begin.
Step F must be finished before step L can begin.
Step Q must be finished before step J can begin.
Step P must be finished before step J can begin.
Step I must be finished before step L can begin.
Step J must be finished before step L can begin.
Step L must be finished before step N can begin.
Step H must be finished before step O can begin.
Step N must be finished before step O can begin.
Step B must be finished before step S can begin.
Step A must be finished before step T can begin.
Step G must be finished before step K can begin.
Step Z must be finished before step N can begin.
Step V must be finished before step I can begin.
Step Z must be finished before step Q can begin.
Step I must be finished before step J can begin.
Step S must be finished before step I can begin.
Step P must be finished before step I can begin.
Step B must be finished before step C can begin.
Step M must be finished before step L can begin.
Step G must be finished before step Z can begin.
Step M must be finished before step C can begin.
Step U must be finished before step F can begin.
Step B must be finished before step Y can begin.
Step W must be finished before step U can begin.
Step G must be finished before step M can begin.
Step M must be finished before step J can begin.
Step C must be finished before step L can begin.
Step K must be finished before step D can begin.
Step S must be finished before step X can begin.
Step Q must be finished before step N can begin.
Step V must be finished before step N can begin.
Step R must be finished before step C can begin.
Step E must be finished before step H can begin.
Step D must be finished before step P can begin.
Step H must be finished before step N can begin.
Step X must be finished before step O can begin.
Step K must be finished before step Y can begin.
Step R must be finished before step F can begin.
Step L must be finished before step O can begin.
Step Y must be finished before step M can begin.
Step T must be finished before step I can begin.
Step T must be finished before step Q can begin.
Step B must be finished before step F can begin.
Step C must be finished before step N can begin.
Step V must be finished before step M can begin.
Step T must be finished before step N can begin.
Step S must be finished before step L can begin.
Step P must be finished before step H can begin.
Step X must be finished before step Q can begin.
Step Z must be finished before step I can begin.
Step Q must be finished before step O can begin.
Step I must be finished before step N can begin.
Step E must be finished before step P can begin.
Step R must be finished before step L can begin.
Step P must be finished before step L can begin.
Step T must be finished before step H can begin.
Step G must be finished before step X can begin.
Step J must be finished before step H can begin.
Step G must be finished before step V can begin.
Step K must be finished before step N can begin.
Step R must be finished before step Q can begin.
Step Z must be finished before step T can begin.
Step E must be finished before step F can begin.
Step Y must be finished before step H can begin.
Step P must be finished before step N can begin.
Step S must be finished before step O can begin.
Step L must be finished before step H can begin.
Step W must be finished before step E can begin.
Step X must be finished before step N can begin.
Step Z must be finished before step D can begin.
Step A must be finished before step H can begin.
Step T must be finished before step X can begin.
Step E must be finished before step Q can begin.
Step K must be finished before step U can begin.
Step M must be finished before step T can begin.
Step J must be finished before step O can begin.
Step D must be finished before step N can begin.
Step K must be finished before step A can begin.
Step G must be finished before step E can begin.
Step R must be finished before step H can begin.
Step W must be finished before step M can begin.
Step U must be finished before step N can begin.
Step Q must be finished before step H can begin.
Step Y must be finished before step A can begin.`

const unorderedGraph = {}
const starts = []
const ends = []
input.split('\n').forEach(line => {
  const start = line.match(/(?<=Step )\w/)[0]
  const end = line.match(/(?<=step )\w/)[0]
  starts.push(start)
  ends.push(end)

  if (!unorderedGraph[end])
    unorderedGraph[end] = []
  unorderedGraph[end].push(start)
})
const startPoints = starts.filter(c => !ends.includes(c)).filter((value, idx, self) => self.indexOf(value) === idx).sort()
startPoints.forEach(x => unorderedGraph[x] = [])

const graph = {};
Object.keys(unorderedGraph).sort().forEach(key => {
  graph[key] = unorderedGraph[key];
})

console.log(graph)


//CABDFE
//

const visited = []
let currentPoint = starts.filter(c => !ends.includes(c)).filter((value, idx, self) => self.indexOf(value) === idx).sort()[0]
for (let i = 0; i < input.split('\n').length; i++) {
  Object.entries(graph).some(([k, v]) => {
    if (v.every(x => visited.includes(x))) {
      currentPoint = k
      delete graph[k]
      return true
    }
  })
  visited.push(currentPoint)
}

console.log(visited.filter((value, idx, self) => self.indexOf(value) === idx).join(''))
