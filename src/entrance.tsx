import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Typography from '@material-ui/core/Typography';
import ReactDOM from 'react-dom';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';

const useStyles = makeStyles((theme) => ({
  header: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}));

function EntranceSection() {
  return (
    <div>
      <h1>Welcome to Minutes Application</h1>
    </div>
  );
}

function EntranceAppBar() {
  const classes = useStyles();

  const toLogin = () => {
    window.location.href = '/login';
  };

  return (
    <div className={classes.header}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            Minutes Application
          </Typography>
          <Button color="inherit" onClick={toLogin}>
            Login
          </Button>
        </Toolbar>
      </AppBar>
    </div>
  );
}

// webpackでバンドルしている関係で存在していないIDが指定される場合がある
// エラーをそのままにしておくと、エラー以後のレンダリングがされない
if (document.getElementById('entrance') != null) {
  ReactDOM.render(<EntranceSection />, document.getElementById('entrance'));
}
if (document.getElementById('entranceHeader') != null) {
  ReactDOM.render(
    <EntranceAppBar />,
    document.getElementById('entranceHeader')
  );
}
