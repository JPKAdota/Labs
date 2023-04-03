const fs = require('fs');

function getSalutation(name) {
  const hour = new Date().getHours();
  let Salutation = '';

  if (hour >= 6 && hour < 12) {
    Salutation = 'Bom dia';
  } else if (hour >= 12 && hour < 18) {
    Salutation = 'Boa tarde';
  } else {
    Salutation = 'Boa noite';
  }

  return `${Salutation}, ${name}!\n`;
}

const names = ['João', 'Igor', 'Richar', 'Symon', 'Tatiana', 'Danilo'];
const filename = 'array_saudacoes_node.txt';

let salutation = '';
for (let i = 0; i < names.length; i++) {
  salutation += getSalutation(names[i]);
}

fs.writeFile(filename, salutation, (err) => {
  if (err) throw err;
  console.log(`Saudações salvas no arquivo ${filename}`);
});
