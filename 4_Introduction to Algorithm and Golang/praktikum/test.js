function isPrime(num) {
  let result;

  if (num === 1) {
    return (result = 'Bukan bilangan prima');
  } else if (num === 2) {
    result = 'Bilangan prima';

    return result;
  } else {
    for (let i = 2; i < num; i++) {
      if (num % i === 0) {
        result = 'Bukan bilangan prima';

        return result;
      } else {
        result = 'Bilangan prima';

        return result;
      }
    }
  }
}

console.log(isPrime(17));

function toggleLamp(num) {
  let toggle = false;
  let status = '';

  for (let i = 1; i <= num; i++) {
    if (num % i === 0) {
      toggle = !toggle;
      if (toggle) {
        status = 'Lampu menyala';
      } else {
        status = 'Lampu mati';
      }
    }
  }

  return status;
}

console.log(toggleLamp(4));
