;; The first three lines of this file were inserted by DrRacket. They record metadata
;; about the language level of this file in a form that our tools can easily process.
#reader(lib "htdp-beginner-reader.ss" "lang")((modname ch1) (read-case-sensitive #t) (teachpacks ()) (htdp-settings #(#t constructor repeating-decimal #f #t none #f ())))
(define (square x) (* x x))
(define (sum-of-squares x y)
  (+ (square x) (square y)))
(define (f a)
  (sum-of-squares (+ a 1) (* a 2)))
(define (abs_s x)
  (if (< x 0)
      (- x)
      x))
(define (>=s x y)
  (or (> x y) (= x y)))
(define a 3)
(define b (+ a 1))
;; 1.3
(define (3ss x y z)
  (+ (square (max x y)) (square (max y z)))
  )
;; 1.5
;;(define (p) (p))
(define (test x y)
  (if (= x 0)
      0
      y))

;; 1.6
(define (sqrt-iter guess x)
  (if (good-enough? guess x)
      guess
      (sqrt-iter (improve guess x)
                 x)))
(define (improve guess x)
  (average guess (/ x guess)))
(define (average x y)
  (/ (+ x y) 2))
(define (good-enough? guess x)
  (< (abs (- (square guess) x)) 0.001))
(define (sqrt_s x)
  (sqrt-iter 1.0 x))

;;1.6
(define (new-if predicate then-clause else-clause)
  (cond (predicate then-clause)
        (else else-clause)))
(define (sqrt-iter_b guess x)
  (new-if (good-enough? guess x)
          guess
          (sqrt-iter_b (improve guess x)
                     x)))
       
(define (cube x)
  (* x (square x) ) )

;;1.7
(define (sqrt-iter2 guess x)
  (* guess x) )
(define (improve2 guess x)
  (* guess x) )
(define (better-enough? previous guess )
    (< (- previous guess) 0.001))
;;    (< (abs (- (square guess) x)) 0.001))
 
;;1.8
(define (cubert-iter guess x)
  (if (good-enough? guess x)
      guess
      (cubert-iter (improve_c guess x)
                 x)))

(define (improve_c guess x)
  (average_c guess (/ x guess)))
(define (average_c x y)
  (/ (+ (/ x (square y) ) (* y 2) ) 3) )
(define (cubert x)
  (cubert-iter 1.0 x))