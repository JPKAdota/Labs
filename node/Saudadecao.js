const readline = require('readline');
const moment = require('moment');
const fs = require('fs');

//Leitura para obter a entrada do usuário
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

// Pergunte ao usuário seu nome
rl.question('Digite seu nome: ', nome => {
  // Obter a hora atual
  const horaAtual = moment().hour();

  // Escolher a saudação apropriada com base na hora atual
  let saudacao;
  if (horaAtual >=6 && horaAtual < 12) {
    saudacao = 'Bom dia';
  } else if (horaAtual >=12 && horaAtual < 18) {
    saudacao = 'Boa tarde';
  } else {
    saudacao = 'Boa noite';
  }

  // Combinar a saudação com o nome fornecido pelo usuário
  const saudacaoCompleta = `${saudacao}, ${nome}`;

  // Escrever a saudação em um arquivo de texto
  fs.writeFile('saudacao.txt', saudacaoCompleta, err => {
    if (err) {
      console.error(err);
      return;
    }
    console.log(`Saudação para ${nome} gravada com sucesso!`);
    rl.close();
  });
});
