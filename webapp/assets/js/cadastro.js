$('#formulario-cadastro').on('submit', criarUsuario)

function criarUsuario(e){
    e.preventDefault()

    if($('#senha').val() != $('#confirmar-senha').val()){
        Swal.fire(
            "Ops...",
            "As senhas não coincidem!",
            "error"
        )
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
        Swal.fire(
            "Sucesso!",
            "Usuário cadastrado com sucesso!",
            "success"
        )
    }).fail(function(response) {
        Swal.fire(
            "Ops...",
            "Erro ao cadastrar usuário!",
            "error"
        )
    })
}