from pydantic import BaseModel
from ..utils.decorators import SchemaRegister


@SchemaRegister.register(
    name='list_directory',
    description='Tool for listing the contents of a directory.'
)
class ListDirectoryTool(BaseModel):
    dirname: str

@SchemaRegister.register(
    name='read_content',
    description='Tool for reading the contents of a file.'
)
class ReadContentTool(BaseModel):
    dirname: str
    filename: str

@SchemaRegister.register(
    name='write_content',
    description='Tool for writing content to a file.'
)
class WriteContentTool(BaseModel):
    dirname: str
    filename: str
    content: str

@SchemaRegister.register(
    name='create_directory',
    description='Tool for creating a new directory.'
)
class CreateDirectoryTool(BaseModel):
    dirpath: str
    dirname: str
