$('#nova-publicacao').on('submit', criarPublicacao)
$(document).on('click', '.curtir-publicacao', curtirPublicacao)
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao)
$('.deletar-publicacao').on('click', deletarPublicacao)

$('#atualizar-publicacao').on('click', atualizarPublicacoes)
$('#descartar-alteracoes').on('click', descartarAlteracoes)

function criarPublicacao(e){
    e.preventDefault()

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(){
        Swal.fire(
            'Sucesso!',
            'Publicação criada com sucesso!',
            'success'
        ).then(function(){
            window.location = "/home"
        })
    }).fail(function(){
        Swal.fire(
            'Ops...',
            'Erro ao criar publicação!',
            'error'
        )
    })
}

function curtirPublicacao(e){
    e.preventDefault()

    const elementoClicado = $(e.target)
    const publicacaoId = elementoClicado.closest('.card')[0].dataset.publicacaoId

    elementoClicado.prop('disabled', true)

    elementoClicado.addClass('descurtir-publicacao')
    elementoClicado.addClass('text-danger')
    elementoClicado.removeClass('curtir-publicacao')

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(){
        const contadorDeCurtidas = elementoClicado.next('span')
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())
        contadorDeCurtidas.text(quantidadeDeCurtidas + 1)
    }).fail(function(){
        Swal.fire(
            'Ops...',
            'Erro ao curtir publicação!',
            'error'
        )
    }).always(function(){
        elementoClicado.prop('disabled', false)
    })
}

function descurtirPublicacao(e){
    e.preventDefault()

    const elementoClicado = $(e.target)
    const publicacaoId = elementoClicado.closest('.card')[0].dataset.publicacaoId

    elementoClicado.prop('disabled', true)

    elementoClicado.removeClass('descurtir-publicacao')
    elementoClicado.removeClass('text-danger')
    elementoClicado.addClass('curtir-publicacao')

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function(){
        const contadorDeCurtidas = elementoClicado.next('span')
        const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text())
        contadorDeCurtidas.text(quantidadeDeCurtidas - 1)
    }).fail(function(){
        Swal.fire(
            'Ops...',
            'Erro ao descurtir publicação!',
            'error'
        )
    }).always(function(){
        elementoClicado.prop('disabled', false)
    })
}

function atualizarPublicacoes(){
    const btn = $(this)
    btn.prop('disabled', true)

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja salvar as alterações?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
       if (!confirmacao.value) return
       
       const publicacaoId = btn[0].dataset.puglicacaoId
   
       $.ajax({
           url: `/publicacoes/${publicacaoId}`,
           method: 'PUT',
           data: {
               titulo: $('#titulo').val(),
               conteudo: $('#conteudo').val(),
           }
       }).done(function(){
           Swal.fire(
               'Sucesso!',
               'Publicação alterada com sucesso!',
               'success'
           ).then(function(){
               window.location = "/home"
           })
       }).fail(function(){
           Swal.fire(
               'Ops...',
               'Erro ao editar publicação!',
               'error'
           )
       }).always(function(){
           btn.prop('disabled', false)
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
       
       window.location = "/home"
    })
}

function deletarPublicacao(e){
    e.preventDefault()

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao){
        if (!confirmacao.value) return

        const elementoClicado = $(e.target)
        const publicacao = elementoClicado.closest('.card')
        const publicacaoId = publicacao[0].dataset.publicacaoId
    
        elementoClicado.prop('disabled', true)
    
        $.ajax({
            url: `/publicacoes/${publicacaoId}`,
            method: "DELETE"
        }).done(function(){
            publicacao.fadeOut("slow", function(){
                $(this).remove()
            })
        }).fail(function(){
            Swal.fire(
                'Ops...',
                'Erro ao deletar publicação!',
                'error'
            )
        }).always(function(){
            elementoClicado.prop('disabled', false)
        })
    })

}