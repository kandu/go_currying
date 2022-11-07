package curry0

// Uncurry to a function of arity-2
func Uncurry2[
  I1 any, I2 any] (
  f func(i1 I1) func(i2 I2)) func(i1 I1, i2 I2) {
    return func(i1 I1, i2 I2) {
      f(i1)(i2) } }

// Uncurry to a function of arity-3
func Uncurry3[
  I1 any, I2 any, I3 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3)) func(i1 I1, i2 I2, i3 I3) {
    return func(i1 I1, i2 I2, i3 I3) {
      f(i1)(i2)(i3) } }

// Uncurry to a function of arity-4
func Uncurry4[
  I1 any, I2 any, I3 any, I4 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4)) func(i1 I1, i2 I2, i3 I3, i4 I4) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4) {
      f(i1) (i2) (i3) (i4) } }

// Uncurry to a function of arity-5
func Uncurry5[
  I1 any, I2 any, I3 any, I4 any, I5 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5) {
      f(i1)(i2) (i3) (i4) (i5) } }

// Uncurry to a function of arity-6
func Uncurry6[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6) {
      f(i1)(i2) (i3) (i4) (i5) (i6) } }

// Uncurry to a function of arity-7
func Uncurry7[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7) } }

// Uncurry to a function of arity-8
func Uncurry8[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8) } }

// Uncurry to a function of arity-9
func Uncurry9[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9) } }

// Uncurry to a function of arity-10
func Uncurry10[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10) } }

// Uncurry to a function of arity-11
func Uncurry11[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11) } }

// Uncurry to a function of arity-12
func Uncurry12[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12) } }

// Uncurry to a function of arity-13
func Uncurry13[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13) } }

// Uncurry to a function of arity-14
func Uncurry14[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14) } }

// Uncurry to a function of arity-15
func Uncurry15[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15) } }

// Uncurry to a function of arity-16
func Uncurry16[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16) } }

// Uncurry to a function of arity-17
func Uncurry17[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17) } }

// Uncurry to a function of arity-18
func Uncurry18[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17)(i18) } }

// Uncurry to a function of arity-19
func Uncurry19[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17)(i18)(i19) } }

// Uncurry to a function of arity-20
func Uncurry20[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17)(i18)(i19)(i20) } }

// Uncurry to a function of arity-21
func Uncurry21[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17)(i18)(i19)(i20)(i21) } }

// Uncurry to a function of arity-22
func Uncurry22[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any, I22 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17)(i18)(i19)(i20)(i21)(i22) } }

// Uncurry to a function of arity-23
func Uncurry23[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17)(i18)(i19)(i20)(i21)(i22)(i23) } }

// Uncurry to a function of arity-24
func Uncurry24[
  I1 any, I2 any, I3 any, I4 any, I5 any, I6 any, I7 any, I8 any, I9 any, I10 any, I11 any, I12 any, I13 any, I14 any, I15 any, I16 any, I17 any, I18 any, I19 any, I20 any, I21 any, I22 any, I23 any, I24 any] (
  f func(i1 I1) func(i2 I2) func(i3 I3) func(i4 I4) func(i5 I5) func(i6 I6) func(i7 I7) func(i7 I8) func(i9 I9) func(i10 I10) func(i11 I11) func(i12 I12) func(i13 I13) func(i14 I14) func(i15 I15) func(i16 I16) func(i17 I17) func(i18 I18) func(i19 I19) func(i20 I20) func(i21 I21) func(i22 I22) func(i23 I23) func(i24 I24)) func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) {
    return func(i1 I1, i2 I2, i3 I3, i4 I4, i5 I5, i6 I6, i7 I7, i8 I8, i9 I9, i10 I10, i11 I11, i12 I12, i13 I13, i14 I14, i15 I15, i16 I16, i17 I17, i18 I18, i19 I19, i20 I20, i21 I21, i22 I22, i23 I23, i24 I24) {
      f(i1)(i2)(i3)(i4)(i5)(i6)(i7)(i8)(i9)(i10)(i11)(i12)(i13)(i14)(i15)(i16)(i17)(i18)(i19)(i20)(i21)(i22)(i23)(i24) } }

