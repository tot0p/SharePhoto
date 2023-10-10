
// Utilisation de la bibliothèque fingerprints2 pour générer une empreinte de navigateur
Fingerprint2.get(function (components) {
    var fingerprint = Fingerprint2.x64hash128(components.map(function (pair) {
        return pair.value;
    }).join(), 31);

    // envoi de l'empreinte au serveur
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/fingerprint', true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({fingerprint: fingerprint}));
});
