FROM python:3.9.7

RUN mkdir -p /workspace
WORKDIR /workspace

# source code
COPY . .

# library install
RUN pip install -r requirements.txt

EXPOSE 64230

CMD python -u ./main.py