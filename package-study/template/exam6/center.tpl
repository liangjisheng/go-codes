{{define "center"}}
<hr>
</div>

<div id="center">

<div id="left">
<hr>
<ul>
<li>Name:  <p>
<li>Id :<p>
<li>Country:
</ul>
<hr>
</div>

<div id="right">
<hr>
<ul>
<li>{{.UserName}}<p>
<li>{{.ID}}<p>
<li>{{.Country}}
</ul>

</div>
</div>
{{end}}