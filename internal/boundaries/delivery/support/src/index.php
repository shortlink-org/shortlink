<?php

phpinfo();

require_once 'config/init.php';

echo msgfmt_format_message('en_US', 'Hello, {name}!', ['name' => 'World']);
