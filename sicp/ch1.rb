def improve(guess, x)
  av = average(guess, (x / guess))
  puts av
  av
end

def good_enough?(guess, x)
  ((guess ** 2) - x).abs < 0.001
end

def sqrt_iter(guess, x)
  if(good_enough?(guess, x))
    guess
  else
    sqrt_iter(improve(guess, x), x)
  end
end

def sqrt_s(x)
  sqrt_iter(1.0, x)
end
