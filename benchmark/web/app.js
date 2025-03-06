const jsFibonacci = (n) => {
  if (n <= 1) {
    return n;
  }

  return jsFibonacci(n - 1) + jsFibonacci(n - 2);
};

const tailFibonacci = (n, left, right) => {
  if (n === 0) {
    return left;
  }

  return tailFibonacci(n - 1, right, left + right);
};

const jsTailFibonacci = (n) => {
  return tailFibonacci(n, 0, 1);
};
