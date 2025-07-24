$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(e){
    e.preventDefault()

    if($('#senha').val() != $('#confirmar-senha').val()){
        alert("As senhas não coincidem!")
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        dataType: "json",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(retorno) {
        alert("Usuário cadastrado com sucesso!")
    }).fail(function(response) {
        alert(response.responseJSON.erro)
    })
}