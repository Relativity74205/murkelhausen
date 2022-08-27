
- Get Port from Electrolama Device: 
    ```bash
    sudo dmesg | grep ttyUSB0
    ```
- Test if flash was successful, after setting the port in the python file
    ```bash
    python znp-uart-test.py
    ```
  
- Flash instructions: https://electrolama.com/radio-docs/flash-cc-bsl/

```bash
python3 cc2538-bsl.py -p /dev/ttyUSB0 -evw CC2652R_coordinator_20220219.hex
```