package curry

// Curry a function of arity-2
func Curry2[
  O any,
  I1 any, I2 any] (
  f func(i1 I1, i2 I2) O) func(i1 I1) func(i2 I2) O {
      return func(i1 I1) func(i2 I2) O {
          return func(i2 I2) O {
              return f(i1, i2) } } }

// Curry a function of arity-3
func Curry3[
  O any,
  I1 any, I2 any, I3 any] (
  f func(i1 I1, i2 I2, i3 I3) O) func(i1 I1) func(i2 I2) func(i3 I3) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) O {
          return func(i2 I2) func(i3 I3) O {
              return func(i3 I3) O {
                  return f(i1, i2, i3) } } } }

// Curry a function of arity-4
func Curry4[
  O any,
  I1 any, I2 any, I3 any, I4 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4) O) func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) O  {
              return func(i3 I3) func(i4 I4) O {
                  return func(i4 I4) O {
                      return f(i1, i2, i3, i4) } } } } }


// Curry a function of arity-5
func Curry5[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) O {
                  return func(i4 I4) func(i5 I5) O {
                      return func(i5 I5) O {
                          return f(i1, i2, i3, i4, i5) } } } } } }

// Curry a function of arity-6
func Curry6[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) O {
                      return func(i5 I5) func(i6 I6) O {
                          return func(i6 I6) O {
                              return f(i1, i2, i3, i4, i5, i6) } } } } } } }

// Curry a function of arity-7
func Curry7[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) O {
                          return func(i6 I6) func(i7 I7) O {
                              return func(i7 I7) O {
                                  return f(i1, i2, i3, i4, i5, i6, i7) } } } } } } } }

// Curry a function of arity-8
func Curry8[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) O {
                              return func(i7 I7) func(i8 I8) O {
                                  return func(i8 I8) O {
                                      return f(i1, i2, i3, i4, i5, i6, i7, i8) } } } } } } } } }

// Curry a function of arity-9
func Curry9[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) O {
                                  return func(i8 I8) func(i9 I9) O {
                                      return func(i9 I9) O {
                                          return f(i1, i2, i3, i4, i5, i6, i7, i8, i9) } } } } } } } } } }

// Curry a function of arity-10
func Curry10[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) O {
                                      return func(i9 I9) func(i10 I10) O {
                                          return func(i10 I10) O {
                                              return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10) } } } } } } } } } } }

// Curry a function of arity-11
func Curry11[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) O {
                                          return func(i10 I10) func(i11 I11) O {
                                              return func(i11 I11) O {
                                                  return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11) } } } } } } } } } } } }

// Curry a function of arity-12
func Curry12[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) O {
                                              return func(i11 I11) func(i12 I12) O {
                                                  return func(i12 I12) O {
                                                      return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12) } } } } } } } } } } } } }

// Curry a function of arity-13
func Curry13[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) O {
                                                  return func(i12 I12) func(i13 I13) O {
                                                      return func(i13 I13) O {
                                                          return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13) } } } } } } } } } } } } } }

// Curry a function of arity-14
func Curry14[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) O {
                                                      return func(i13 I13) func(i14 I14) O {
                                                          return func(i14 I14) O {
                                                              return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14) } } } } } } } } } } } } } } }

// Curry a function of arity-15
func Curry15[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) O {
                                                          return func(i14 I14) func(i15 I15) O {
                                                              return func(i15 I15) O {
                                                                  return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15) } } } } } } } } } } } } } } } }

// Curry a function of arity-16
func Curry16[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) O {
                                                              return func(i15 I15) func(i16 I16) O {
                                                                  return func(i16 I16) O {
                                                                      return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16) } } } } } } } } } } } } } } } } }

// Curry a function of arity-17
func Curry17[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) O {
                                                                  return func(i16 I16) func(i17 I17) O {
                                                                      return func(i17 I17) O {
                                                                          return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17) } } } } } } } } } } } } } } } } } }

// Curry a function of arity-18
func Curry18[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) O {
                                                                  return func(i16 I16) func(i17 I17) func(i18 I18) O {
                                                                      return func(i17 I17) func(i18 I18) O {
                                                                          return func(i18 I18) O {
                                                                              return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18) } } } } } } } } } } } } } } } } } } }

// Curry a function of arity-19
func Curry19[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                                                  return func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) O {
                                                                      return func(i17 I17) func(i18 I18) func(i19 I19) O {
                                                                          return func(i18 I18) func(i19 I19) O {
                                                                              return func(i19 I19) O {
                                                                                  return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19) } } } } } } } } } } } } } } } } } } } }

// Curry a function of arity-20
func Curry20[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                                  return func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                                      return func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                                          return func(i18 I18) func(i19 I19) func(i20 I20) O {
                                                                              return func(i19 I19) func(i20 I20) O {
                                                                                  return func(i20 I20) O {
                                                                                      return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20) } } } } } } } } } } } } } } } } } } } } }

// Curry a function of arity-21
func Curry21[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                                  return func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                                      return func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                                          return func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                                              return func(i19 I19) func(i20 I20) func(i21 I21) O {
                                                                                  return func(i20 I20) func(i21 I21) O {
                                                                                      return func(i21 I21) O {
                                                                                          return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21) } } } } } } } } } } } } } } } } } } } } } }

// Curry a function of arity-22
func Curry22[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any, I22 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                                  return func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                                      return func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                                          return func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                                              return func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                                                  return func(i20 I20) func(i21 I21) func(i22 I22) O {
                                                                                      return func(i21 I21) func(i22 I22) O {
                                                                                          return func(i22 I22) O {
                                                                                              return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22) } } } } } } } } } } } } } } } } } } } } } } }

// Curry a function of arity-23
func Curry23[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                                  return func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                                      return func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                                          return func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                                              return func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                                                  return func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                                                      return func(i21 I21) func(i22 I22) func(i23 I23) O {
                                                                                          return func(i22 I22) func(i23 I23) O {
                                                                                              return func(i23 I23) O {
                                                                                                  return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23) } } } } } } } } } } } } } } } } } } } } } } } }

// Curry a function of arity-24
func Curry24[
  O any,
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any, I24 any] (
  f func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) O) func(i1 I1)func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
      return func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
          return func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O  {
              return func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                  return func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                      return func(i5 I5) func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                          return func(i6 I6) func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                              return func(i7 I7) func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                  return func(i8 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                      return func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                          return func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                              return func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                  return func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                      return func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                          return func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                              return func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                  return func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                      return func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                          return func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                              return func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                                  return func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                                      return func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                                          return func(i22 I22) func(i23 I23) func(i24 I24) O {
                                                                                              return func(i23 I23) func(i24 I24) O {
                                                                                                  return func(i24 I24) O {
                                                                                                      return f(i1, i2, i3, i4, i5, i6, i7, i8, i9, i10, i11, i12, i13, i14, i15, i16, i17, i18, i19, i20, i21, i22, i23, i24) } } } } } } } } } } } } } } } } } } } } } } } } }

