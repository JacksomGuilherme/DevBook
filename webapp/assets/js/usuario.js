$('#parar-de-seguir').on('click', pararDeSeguir)
$('#seguir').on('click', seguir)

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