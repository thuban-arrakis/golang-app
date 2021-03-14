<html>
    <head>
    <link rel="stylesheet" href="./stylesheets/main.css">
    <title></title>
    </head>
    <body>
        <form action="/login">
           <label>Method</label>
             <select class="form-control">
              <option>GET</option>
              <option>POST</option>
              <option>PUT</option>
              <option>DELETE</option>
              <option>UPDATE</option>
           </select>
            Data:<input type="password" name="password">
            <input type="submit" value="Login">
            <p>Длина в дюймах: <output name="result">0</output></p>
        </form>
    </body>
</html>