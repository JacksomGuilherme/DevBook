$('#nova-publicacao').on('submit', criarPublicacao)
$(document).on('click', '.curtir-publicacao', curtirPublicacao)
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao)

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
        window.location = "/home"
    }).fail(function(){
        alert('Erro ao criar publicação!')
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
        alert('Erro ao curtir publicação!')
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
        alert('Erro ao descurtir publicação!')
    }).always(function(){
        elementoClicado.prop('disabled', false)
    })
}