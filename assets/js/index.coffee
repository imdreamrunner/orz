GetIdeas = (callback) ->
    $.ajax
        url: '/ideas'
        success: (ret) ->
            console.log ret
            callback(ret)

DisplayIdeaList = (items) ->
    template = $('#list-item-template').html()
    $list = $('#idea-list')
    console.log $list
    for item in items
        html = template
        html = html.replace /{{ content }}/g, item['content']
        html = html.replace /{{ html }}/g, item['html']
        $item = $(html)
        $item.find('.name').html item['name']
        if item['link']
            $item.find('.name').attr 'href', item['link']
        $item.find('.header-avatar').attr 'src', 'http://www.gravatar.com/avatar/' + md5(item['email'])
        $item.find('.time').html moment(item['timestamp'], 'X').fromNow()
        $list.append($item)
    $('.collapsible').collapsible()

SubmitIdea = ->
    name = $('#name').val()
    email = $('#email').val()
    link = $('#link').val()
    content = $('#content').val()
    postData =
        name: name
        email: email
        link: link
        content: content
    $.ajax
        url: '/post'
        method: 'post'
        data: postData
        success: (ret) ->
            window.location.reload()
        error: (ret) ->
            msg = ret.responseText
            msg = JSON.parse(msg)['message']
            Materialize.toast(msg, 4000)

$(document).ready ->
    console.log "hello world"
    GetIdeas(DisplayIdeaList)
    $('#submit-idea-form').submit ->
        SubmitIdea()
        return false
