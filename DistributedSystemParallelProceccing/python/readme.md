## RMI server-client
install package `Pyro5` by
```bash
pip install Pyro5
```
### RMI server on same machine
- to run rmi_server, first run this command for nameserver:
```bash
python -m Pyro5.nameserver
```
- then run server on another terminal and client on 3rd terminal.

### RMI server on different machine
- to run rmi_server, first run this command for nameserver:
```bash
python -m Pyro5.nameserver --host 0.0.0.0
```
- replace ip address of server on `rmi_server.py` instead of `localhost`
- then run `rmi_server.py` on server machine.
- and run `rmi_client.py` on client machine.
- `rmi_interface` should have on both machine