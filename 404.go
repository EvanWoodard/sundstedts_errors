package httppages

const page404 string = `
{{ define "404" }}
<!DOCTYPE html>
<html>
<head>
    <title>Sundstedt's - 404</title>
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <link href="https://fonts.googleapis.com/css?family=Roboto&display=swap" rel="stylesheet">
    <link href="https://cdn.sundstedt.us/css/theme.light.css" rel="stylesheet">
    <link href="https://cdn.sundstedt.us/css/mobile.css" rel="stylesheet">
    <link href="https://cdn.sundstedt.us/css/standard.css" rel="stylesheet">
</head>
<body>
    <div class="topbar">
        <div class="logo"><img src="https://cdn.sundstedt.us/img/android-chrome-512x512.png" /></div>
        <div class="title">The Sundstedt's</div>
        <div class="spacer"></div>
        <form class="redirect-form" action="https://iam.sundstedt.us/logout" method="POST">
            <input id="home_redirect" type="hidden" name="redirect" value="https://evenson.sundstedt.us/">
            <button class="btn login-btn" type="submit">Home</button>
        </form>
    </div>
    <div class="page-content">
        <div class="http-error">
            <strong class="code">404</strong>
            <div class="message">The page you're looking for cannot be found... Oops!</div>
        </div>
    </div>
</body>
</html>
{{ end }}
`
