$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)
$('#editar-usuario').on('submit', editar)
$('#alterar-senha').on('submit', atualizarSenha)
$('#deletar-usuario').on('click', deletarUsuario)

$('#descartar-alteracoes').on('click', descartarAlteracoes)

function pararDeSeguir(){
    const btn = $(this)
    const usuarioId = btn[0].dataset.usuarioId
    btn.prop('disabled', true)

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function(){
        window.location = `/usuario/${usuarioId}`
    }).fail(function(){
        Swal.fire(
            'Ops...',
            'Erro ao parar de seguir usuário!',
            'error'
        )
        btn.prop('disabled', false)
    })
}

function seguir(){
    const btn = $(this)
    const usuarioId = btn[0].dataset.usuarioId
    btn.prop('disabled', true)

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function(){
        window.location = `/usuario/${usuarioId}`
    }).fail(function(){
        Swal.fire(
            'Ops...',
            'Erro ao seguir usuário!',
            'error'
        )
        btn.prop('disabled', false)
    })
}

function editar(e){
    e.preventDefault()
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja salvar as alterações?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
        if (!confirmacao.value) return

        $.ajax({
            url: "/editar-usuario",
            method: "PUT",
            data: {
                nome: $('#nome').val(),
                email: $('#email').val(),
                nick: $('#nick').val()
            }
        }).done(function(){
            Swal.fire(
                'Sucesso',
                'Usuário alterado com sucesso!',
                'success'
            ).then(function(){
                window.location = "/perfil"
            })
        }).fail(function(){
            Swal.fire(
                'Ops...',
                'Erro ao alterar o usuário!',
                'error'
            )
        })
    })
}

function descartarAlteracoes(){
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja descartar as alterações?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
       if (!confirmacao.value) return
       
       window.location = "/perfil"
    })
}

function atualizarSenha(e){
    e.preventDefault()

    if ($('#nova-senha').val() != $('#confirmar-senha').val()){
        Swal.fire(
            'Ops...',
            'As senhas não coincidem!',
            'warning'
        )
        return
    }

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja alterar a senha?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
       if (!confirmacao.value) return
       
       $.ajax({
           url: "/alterar-senha",
           method: "POST",
           data: {
               atual: $('#senha-atual').val(),
               nova: $('#nova-senha').val(), 
           }
       }).done(function(){
           Swal.fire(
               'Sucesso',
               'Senha alterada com sucesso!',
               'success'
           ).then(function(){
               window.location = "/perfil"
           })
       }).fail(function(){
           Swal.fire(
               'Ops...',
               'Erro ao alterar a senha!',
               'error'
           )
       })
    })
}

function deletarUsuario(){
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir a sua conta?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",

        icon: "warning"
    }).then(function(confirmacao){
       if (!confirmacao.value) return
       
       $.ajax({
           url: "/deletar-usuario",
           method: "DELETE",
       }).done(function(){
           Swal.fire(
               'Sucesso',
               'Conta excluída com sucesso!',
               'success'
           ).then(function(){
               window.location = "/logout"
           })
       }).fail(function(){
           Swal.fire(
               'Ops...',
               'Erro ao excluir a conta!',
               'error'
           )
       })
    })
}
