$('#login').on('submit', fazerLogin)

function fazerLogin(e) {
    e.preventDefault()

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#email').val(),
            senha: $('#senha').val()
        }
    }).done(function(retorno) {
        window.location = "/home"
    }).fail(function(response) {
        alert("Usuário ou senha inválidos!")
    })
}
