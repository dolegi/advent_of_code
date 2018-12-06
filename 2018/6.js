let input = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

input = `192, 220
91, 338
65, 319
143, 310
243, 205
237, 135
342, 197
114, 56
189, 168
194, 174
55, 331
181, 162
272, 111
201, 121
73, 88
276, 274
324, 323
201, 146
125, 301
190, 185
247, 307
157, 65
217, 181
62, 222
319, 202
86, 342
333, 339
181, 240
263, 198
200, 296
306, 228
55, 50
154, 356
54, 70
91, 91
265, 182
272, 267
118, 296
75, 140
319, 272
357, 341
193, 342
102, 334
246, 123
328, 139
229, 284
199, 309
151, 243
295, 229
201, 277`

const matrix = []
let biggestY = 800
let biggestX = 800
const coords = []

input.split('\n').forEach((coord, idx) => {
  const [x, y] = coord.split(',').map(x => parseInt(x) - 1)

  if (x > biggestX)
    biggestX = x

  if (y > biggestY)
    biggestY = y
})

for (let y = 0; y <= biggestY; y++) {
  if (!matrix[y]) {
    matrix[y] = []
  }
  for (let x = 0; x <= biggestX; x++) {
    matrix[y][x] = null
  }
}

input.split('\n').forEach((coord, idx) => {
  const [x, y] = coord.split(',').map(x => parseInt(x) - 1)

  if (idx < 26) {
    matrix[y][x] = String.fromCharCode(idx + 65)
  } else {
    matrix[y][x] = String.fromCharCode(idx + 65 - 25) + String.fromCharCode(idx + 65 - 25)
  }
  coords.push([x, y])
})

matrix.forEach((row, y) => {
  row.forEach((cel, x) => {
    if (cel !== null)
      return

    let closest = Infinity
    coords.forEach(([x2, y2], idx) => {
      const dist = manDistance(x, x2, y, y2)
      if (dist === closest) {
        matrix[y][x] = '.'
      } else if (dist < closest) {
        closest = dist
        if (idx < 26) {
          matrix[y][x] = String.fromCharCode(idx + 97)
        } else {
          matrix[y][x] = String.fromCharCode(idx + 97 - 25) + String.fromCharCode(idx + 97 - 25)
        }
      }
    })
  })
})

function manDistance(a0, a1, b0, b1) {
  return Math.abs(a0 - a1) + Math.abs(b0 - b1)
}

matrix.forEach((row, y) => {
  row.forEach((cel, x) => {
    matrix[y][x] = matrix[y][x].toLowerCase()
  })
})

let invalid = []
matrix[0].forEach(cel => {
  invalid.push(cel)
})

matrix.forEach(row => {
  const cel = row[0]
  invalid.push(cel)
})

matrix[matrix.length - 1].forEach(cel => {
  invalid.push(cel)
})

matrix.forEach(row => {
  const cel = row[row.length - 1]
  invalid.push(cel)
})

matrix.forEach((row, y) => {
  // matrix[y] = row.filter(cel => !invalid.includes(cel))
})

matrix.forEach((row, y) => {
  row.forEach((cel, x) => {
    matrix[y][x] = matrix[y][x].toLowerCase()
  })
})

const count = {}
matrix.forEach((row, y) => {
  row.forEach((cel, x) => {
    if (!count[cel])
      count[cel] = 0

    count[cel] += 1
  })
})

console.log(Math.max(...Object.values(count)))

matrix.forEach((row, y) => {
  row.forEach((cel, x) => {
    matrix[y][x] = 0
    coords.forEach(([x2, y2], idx) => {
      const dist = manDistance(x, x2, y, y2)
      matrix[y][x] += dist
    })
  })
})

matrix.forEach((row, y) => {
  matrix[y] = row.filter(cel => cel < 10000)
})

// matrix.forEach((row, y) => {
//   console.log(row)
// })

let c = 0
matrix.forEach((row, y) => {
  c += row.length
})

console.log(c)

