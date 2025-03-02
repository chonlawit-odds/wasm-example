const jsFibonacci = (n) => {
  if (n <= 1) {
    return n;
  }

  return jsFibonacci(n - 1) + jsFibonacci(n - 2);
};
