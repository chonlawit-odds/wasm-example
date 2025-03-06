const initFibonacci = () => {
  document.getElementById("statusFibonacci").textContent = ``;
  document.getElementById("runButtonFibonacci").disabled = true;
  document.getElementById("numberFibonacci").disabled = true;
};

const finishFibonacci = () => {
  document.getElementById("statusFibonacci").textContent = `Done`;
  document.getElementById("runButtonFibonacci").disabled = false;
  document.getElementById("numberFibonacci").disabled = false;
};

const runBenchmark = (fn, n, loop = 5) => {
  return new Promise((resolve, reject) => {
    console.log("Running test with n as ", n);

    // Execute function
    const res = fn(n);
    console.log("Result:", res);

    let elapsedTime = 0.0;
    for (let lap = 0; lap < loop; lap++) {
      const startTime = performance.now();
      fn(n);
      const endTime = performance.now();
      const diff = endTime - startTime;
      elapsedTime += diff;
      console.log(lap, diff);
    }

    // Average
    resolve((elapsedTime / loop).toFixed(6));
  });
};

const executeBenchmarkFibonacci = (arr, index, n) => {
  if (index >= arr.length) {
    finishFibonacci();
    return;
  }

  const benchmark = arr[index];
  document.getElementById(
    "statusFibonacci"
  ).textContent = `Running ${benchmark.name}...`;
  setTimeout(async () => {
    const res = await runBenchmark(benchmark.func, n);
    document.getElementById(benchmark.elem).textContent = res;
    setTimeout(() => executeBenchmarkFibonacci(arr, index + 1, n));
  });
};

const startFibonacci = () => {
  console.log("Starting benchmark fibonacci");

  console.log(window);

  const availableTests = [
    ////////// Fibonacci recursive
    // { elem: "jsPerformanceFibonacci", func: jsFibonacci, name: "JavaScript" },
    // { elem: "goPerformanceFibonacci", func: wsGoFibonacci, name: "Go" },
    // {
    //   elem: "goRoutinePerformanceFibonacci",
    //   func: wsGoRoutineFibonacci,
    //   name: "Go (Routine)",
    // },
    ////////// Fibonacci tail recursive
    // {
    //   elem: "jsPerformanceFibonacci",
    //   func: jsTailFibonacci,
    //   name: "JavaScript",
    // },
    // { elem: "goPerformanceFibonacci", func: wsGoTailFibonacci, name: "Go" },
    // {
    //   elem: "goRoutinePerformanceFibonacci",
    //   func: wsGoRoutineTailFibonacci,
    //   name: "Go (Routine)",
    // },
    ////////// TinyGo
    { elem: "jsPerformanceFibonacci", func: jsFibonacci, name: "JavaScript" },
    {
      elem: "tinygoPerformanceFibonacci",
      func: wasm.exports.wsTinyFibonacci,
      name: "TinyGo",
    },
  ];
  initFibonacci();

  availableTests.forEach(
    (el) => (document.getElementById(el.elem).textContent = "")
  );
  const inputNumber = document.getElementById("numberFibonacci").value;
  const n = Math.floor(parseFloat(inputNumber));
  setTimeout(() => executeBenchmarkFibonacci(availableTests, 0, n));
};
